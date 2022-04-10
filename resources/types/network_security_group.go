package types

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/network_security_group"
	rg "github.com/multycloud/multy/resources/resource_group"
	"github.com/multycloud/multy/validate"
	"strconv"
	"strings"
)

/*
Notes
NSG can only be applied to NIC (currently done in VM creation, to be changed to separate resource)
When NSG is applied, only rules specified are allowed.
AWS: VPC traffic is always added as an extra rule
*/

type NetworkSecurityGroup struct {
	resources.ResourceWithId[*resourcespb.NetworkSecurityGroupArgs]

	VirtualNetwork *VirtualNetwork
}

type RuleType struct {
	Protocol  string `cty:"protocol"`
	Priority  int    `cty:"priority"`
	FromPort  string `cty:"from_port"`
	ToPort    string `cty:"to_port"`
	CidrBlock string `cty:"cidr_block"`
	Direction string `cty:"direction"`
}

const (
	INGRESS = "ingress"
	EGRESS  = "egress"
	BOTH    = "both"
	ALLOW   = "allow"
	DENY    = "deny"
)

func NewNetworkSecurityGroup(resourceId string, args *resourcespb.NetworkSecurityGroupArgs, others resources.Resources) (*NetworkSecurityGroup, error) {
	vn, err := resources.Get[*VirtualNetwork](resourceId, others, args.VirtualNetworkId)
	if err != nil {
		return nil, err
	}
	return &NetworkSecurityGroup{
		ResourceWithId: resources.ResourceWithId[*resourcespb.NetworkSecurityGroupArgs]{
			ResourceId: resourceId,
			Args:       args,
		},
		VirtualNetwork: vn,
	}, nil
}

func (r *NetworkSecurityGroup) Translate(resources.MultyContext) ([]output.TfBlock, error) {
	if r.GetCloud() == commonpb.CloudProvider_AWS {
		awsRules := translateAwsNsgRules(r.Args.Rules)

		allowVpcTraffic := network_security_group.AwsSecurityGroupRule{
			Protocol:   "-1",
			FromPort:   0,
			ToPort:     0,
			CidrBlocks: []string{r.VirtualNetwork.Args.CidrBlock},
		}

		awsRules[INGRESS] = append(awsRules[INGRESS], allowVpcTraffic)
		awsRules[EGRESS] = append(awsRules[EGRESS], allowVpcTraffic)

		vnId, err := resources.GetMainOutputId(r.VirtualNetwork)
		if err != nil {
			return nil, err
		}
		return []output.TfBlock{
			network_security_group.AwsSecurityGroup{
				AwsResource: common.NewAwsResource(r.ResourceId, r.Args.Name),
				VpcId:       vnId,
				Name:        r.Args.Name,
				Description: "Managed by Multy",
				Ingress:     awsRules["ingress"],
				Egress:      awsRules["egress"],
			},
		}, nil
	} else if r.GetCloud() == commonpb.CloudProvider_AZURE {
		return []output.TfBlock{
			network_security_group.AzureNsg{
				AzResource: common.NewAzResource(
					r.ResourceId, r.Args.Name, rg.GetResourceGroupName(r.Args.CommonParameters.ResourceGroupId),
					r.GetCloudSpecificLocation(),
				),
				Rules: translateAzNsgRules(r.Args.Rules),
			},
		}, nil
	}
	return nil, fmt.Errorf("cloud %s is not supported for this resource type ", r.GetCloud().String())
}

func translateAwsNsgRules(rules []*resourcespb.NetworkSecurityRule) map[string][]network_security_group.AwsSecurityGroupRule {
	awsRules := map[string][]network_security_group.AwsSecurityGroupRule{}

	for _, rule := range rules {
		awsFromPort := int(rule.PortRange.From)
		awsToPort := int(rule.PortRange.To)

		awsProtocol := rule.Protocol
		if rule.Protocol == "*" {
			awsProtocol = "-1"
			awsFromPort = 0
			awsToPort = 0
		}

		if rule.Direction == resourcespb.Direction_BOTH_DIRECTIONS {
			awsRules[INGRESS] = append(
				awsRules[INGRESS], network_security_group.AwsSecurityGroupRule{
					Protocol:   awsProtocol,
					FromPort:   awsFromPort,
					ToPort:     awsToPort,
					CidrBlocks: []string{rule.CidrBlock},
				},
			)
			awsRules[EGRESS] = append(
				awsRules[EGRESS], network_security_group.AwsSecurityGroupRule{
					Protocol:   awsProtocol,
					FromPort:   awsFromPort,
					ToPort:     awsToPort,
					CidrBlocks: []string{rule.CidrBlock},
				},
			)
		} else {
			awsRules[BOTH] = append(
				awsRules[BOTH], network_security_group.AwsSecurityGroupRule{
					Protocol:   awsProtocol,
					FromPort:   awsFromPort,
					ToPort:     awsToPort,
					CidrBlocks: []string{rule.CidrBlock},
				},
			)
		}
	}
	return awsRules
}

func translatePortRange(pr *resourcespb.PortRange) string {
	from := "*"
	if pr.From != 0 {
		from = strconv.Itoa(int(pr.From))
	}
	to := "*"
	if pr.To != 0 {
		to = strconv.Itoa(int(pr.To))
	}
	return fmt.Sprintf("%s-%s", from, to)
}

func translateAzNsgRules(rules []*resourcespb.NetworkSecurityRule) []network_security_group.AzureRule {
	m := map[resourcespb.Direction]string{
		resourcespb.Direction_INGRESS: "Inbound",
		resourcespb.Direction_EGRESS:  "Outbound",
	}

	var rls []network_security_group.AzureRule

	for _, rule := range rules {
		protocol := strings.Title(strings.ToLower(rule.Protocol))
		if rule.Direction == resourcespb.Direction_BOTH_DIRECTIONS {
			rls = append(
				rls, network_security_group.AzureRule{
					Name:                     strconv.Itoa(len(rls)),
					Protocol:                 protocol,
					Priority:                 int(rule.Priority),
					Access:                   "Allow",
					SourcePortRange:          "*",
					SourceAddressPrefix:      "*",
					DestinationPortRange:     translatePortRange(rule.PortRange),
					DestinationAddressPrefix: "*",
					Direction:                m[resourcespb.Direction_INGRESS],
				},
			)
			rls = append(
				rls, network_security_group.AzureRule{
					Name:                     strconv.Itoa(len(rls)),
					Protocol:                 protocol,
					Priority:                 int(rule.Priority),
					Access:                   "Allow",
					SourcePortRange:          "*",
					SourceAddressPrefix:      "*",
					DestinationPortRange:     translatePortRange(rule.PortRange),
					DestinationAddressPrefix: "*",
					Direction:                m[resourcespb.Direction_EGRESS],
				},
			)
		} else {
			rls = append(
				rls, network_security_group.AzureRule{
					Name:                     strconv.Itoa(len(rls)),
					Protocol:                 protocol,
					Priority:                 int(rule.Priority),
					Access:                   "Allow",
					SourcePortRange:          "*",
					SourceAddressPrefix:      "*",
					DestinationPortRange:     translatePortRange(rule.PortRange),
					DestinationAddressPrefix: "*",
					Direction:                m[rule.Direction],
				},
			)
		}
	}

	return rls
}

func validateRuleDirection(s string) bool {
	return s == INGRESS || s == EGRESS || s == BOTH
}

func validateRuleAction(s string) bool {
	return s == ALLOW || s == DENY
}

func validatePort(port int32) bool {
	return port >= 0 && port <= 65535
}

func (r *NetworkSecurityGroup) Validate(ctx resources.MultyContext) (errs []validate.ValidationError) {
	errs = append(errs, r.ResourceWithId.Validate()...)
	for _, rule := range r.Args.Rules {
		if !validatePort(rule.PortRange.To) {
			errs = append(errs, r.NewValidationError(fmt.Sprintf("rule to_port \"%d\" is not valid", rule.PortRange.To), "rules"))
		}
		if !validatePort(rule.PortRange.From) {
			errs = append(errs, r.NewValidationError(fmt.Sprintf("rule from_port \"%d\" is not valid", rule.PortRange.From), "rules"))
		}
		// TODO validate CIDR
		//		validate protocol
	}
	// TODO validate location matches with VN location
	return errs
}

func (r *NetworkSecurityGroup) GetMainResourceName() (string, error) {
	switch r.GetCloud() {
	case commonpb.CloudProvider_AWS:
		return network_security_group.AwsSecurityGroupResourceName, nil
	case commonpb.CloudProvider_AZURE:
		return network_security_group.AzureNetworkSecurityGroupResourceName, nil
	default:
		return "", fmt.Errorf("unknown cloud %s", r.GetCloud().String())
	}
}

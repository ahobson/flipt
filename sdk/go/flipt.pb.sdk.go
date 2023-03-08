// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.

package sdk

import (
	context "context"
	flipt "go.flipt.io/flipt/rpc/flipt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Flipt struct {
	transport FliptTransport
}

type FliptTransport interface {
	Evaluate(context.Context, *flipt.EvaluationRequest) (*flipt.EvaluationResponse, error)
	BatchEvaluate(context.Context, *flipt.BatchEvaluationRequest) (*flipt.BatchEvaluationResponse, error)
	GetFlag(context.Context, *flipt.GetFlagRequest) (*flipt.Flag, error)
	ListFlags(context.Context, *flipt.ListFlagRequest) (*flipt.FlagList, error)
	CreateFlag(context.Context, *flipt.CreateFlagRequest) (*flipt.Flag, error)
	UpdateFlag(context.Context, *flipt.UpdateFlagRequest) (*flipt.Flag, error)
	DeleteFlag(context.Context, *flipt.DeleteFlagRequest) (*emptypb.Empty, error)
	CreateVariant(context.Context, *flipt.CreateVariantRequest) (*flipt.Variant, error)
	UpdateVariant(context.Context, *flipt.UpdateVariantRequest) (*flipt.Variant, error)
	DeleteVariant(context.Context, *flipt.DeleteVariantRequest) (*emptypb.Empty, error)
	GetRule(context.Context, *flipt.GetRuleRequest) (*flipt.Rule, error)
	ListRules(context.Context, *flipt.ListRuleRequest) (*flipt.RuleList, error)
	CreateRule(context.Context, *flipt.CreateRuleRequest) (*flipt.Rule, error)
	UpdateRule(context.Context, *flipt.UpdateRuleRequest) (*flipt.Rule, error)
	OrderRules(context.Context, *flipt.OrderRulesRequest) (*emptypb.Empty, error)
	DeleteRule(context.Context, *flipt.DeleteRuleRequest) (*emptypb.Empty, error)
	CreateDistribution(context.Context, *flipt.CreateDistributionRequest) (*flipt.Distribution, error)
	UpdateDistribution(context.Context, *flipt.UpdateDistributionRequest) (*flipt.Distribution, error)
	DeleteDistribution(context.Context, *flipt.DeleteDistributionRequest) (*emptypb.Empty, error)
	GetSegment(context.Context, *flipt.GetSegmentRequest) (*flipt.Segment, error)
	ListSegments(context.Context, *flipt.ListSegmentRequest) (*flipt.SegmentList, error)
	CreateSegment(context.Context, *flipt.CreateSegmentRequest) (*flipt.Segment, error)
	UpdateSegment(context.Context, *flipt.UpdateSegmentRequest) (*flipt.Segment, error)
	DeleteSegment(context.Context, *flipt.DeleteSegmentRequest) (*emptypb.Empty, error)
	CreateConstraint(context.Context, *flipt.CreateConstraintRequest) (*flipt.Constraint, error)
	UpdateConstraint(context.Context, *flipt.UpdateConstraintRequest) (*flipt.Constraint, error)
	DeleteConstraint(context.Context, *flipt.DeleteConstraintRequest) (*emptypb.Empty, error)
}

func (x *Flipt) Evaluate(ctx context.Context, v *flipt.EvaluationRequest) (*flipt.EvaluationResponse, error) {
	return x.transport.Evaluate(ctx, v)
}

func (x *Flipt) BatchEvaluate(ctx context.Context, v *flipt.BatchEvaluationRequest) (*flipt.BatchEvaluationResponse, error) {
	return x.transport.BatchEvaluate(ctx, v)
}

func (x *Flipt) GetFlag(ctx context.Context, key string) (*flipt.Flag, error) {
	return x.transport.GetFlag(ctx, &flipt.GetFlagRequest{Key: key})
}

func (x *Flipt) ListFlags(ctx context.Context, v *flipt.ListFlagRequest) (*flipt.FlagList, error) {
	return x.transport.ListFlags(ctx, v)
}

func (x *Flipt) CreateFlag(ctx context.Context, v *flipt.CreateFlagRequest) (*flipt.Flag, error) {
	return x.transport.CreateFlag(ctx, v)
}

func (x *Flipt) UpdateFlag(ctx context.Context, v *flipt.UpdateFlagRequest) (*flipt.Flag, error) {
	return x.transport.UpdateFlag(ctx, v)
}

func (x *Flipt) DeleteFlag(ctx context.Context, key string) error {
	_, err := x.transport.DeleteFlag(ctx, &flipt.DeleteFlagRequest{Key: key})
	return err
}

func (x *Flipt) CreateVariant(ctx context.Context, v *flipt.CreateVariantRequest) (*flipt.Variant, error) {
	return x.transport.CreateVariant(ctx, v)
}

func (x *Flipt) UpdateVariant(ctx context.Context, v *flipt.UpdateVariantRequest) (*flipt.Variant, error) {
	return x.transport.UpdateVariant(ctx, v)
}

func (x *Flipt) DeleteVariant(ctx context.Context, v *flipt.DeleteVariantRequest) error {
	_, err := x.transport.DeleteVariant(ctx, v)
	return err
}

func (x *Flipt) GetRule(ctx context.Context, v *flipt.GetRuleRequest) (*flipt.Rule, error) {
	return x.transport.GetRule(ctx, v)
}

func (x *Flipt) ListRules(ctx context.Context, v *flipt.ListRuleRequest) (*flipt.RuleList, error) {
	return x.transport.ListRules(ctx, v)
}

func (x *Flipt) CreateRule(ctx context.Context, v *flipt.CreateRuleRequest) (*flipt.Rule, error) {
	return x.transport.CreateRule(ctx, v)
}

func (x *Flipt) UpdateRule(ctx context.Context, v *flipt.UpdateRuleRequest) (*flipt.Rule, error) {
	return x.transport.UpdateRule(ctx, v)
}

func (x *Flipt) OrderRules(ctx context.Context, v *flipt.OrderRulesRequest) error {
	_, err := x.transport.OrderRules(ctx, v)
	return err
}

func (x *Flipt) DeleteRule(ctx context.Context, v *flipt.DeleteRuleRequest) error {
	_, err := x.transport.DeleteRule(ctx, v)
	return err
}

func (x *Flipt) CreateDistribution(ctx context.Context, v *flipt.CreateDistributionRequest) (*flipt.Distribution, error) {
	return x.transport.CreateDistribution(ctx, v)
}

func (x *Flipt) UpdateDistribution(ctx context.Context, v *flipt.UpdateDistributionRequest) (*flipt.Distribution, error) {
	return x.transport.UpdateDistribution(ctx, v)
}

func (x *Flipt) DeleteDistribution(ctx context.Context, v *flipt.DeleteDistributionRequest) error {
	_, err := x.transport.DeleteDistribution(ctx, v)
	return err
}

func (x *Flipt) GetSegment(ctx context.Context, key string) (*flipt.Segment, error) {
	return x.transport.GetSegment(ctx, &flipt.GetSegmentRequest{Key: key})
}

func (x *Flipt) ListSegments(ctx context.Context, v *flipt.ListSegmentRequest) (*flipt.SegmentList, error) {
	return x.transport.ListSegments(ctx, v)
}

func (x *Flipt) CreateSegment(ctx context.Context, v *flipt.CreateSegmentRequest) (*flipt.Segment, error) {
	return x.transport.CreateSegment(ctx, v)
}

func (x *Flipt) UpdateSegment(ctx context.Context, v *flipt.UpdateSegmentRequest) (*flipt.Segment, error) {
	return x.transport.UpdateSegment(ctx, v)
}

func (x *Flipt) DeleteSegment(ctx context.Context, key string) error {
	_, err := x.transport.DeleteSegment(ctx, &flipt.DeleteSegmentRequest{Key: key})
	return err
}

func (x *Flipt) CreateConstraint(ctx context.Context, v *flipt.CreateConstraintRequest) (*flipt.Constraint, error) {
	return x.transport.CreateConstraint(ctx, v)
}

func (x *Flipt) UpdateConstraint(ctx context.Context, v *flipt.UpdateConstraintRequest) (*flipt.Constraint, error) {
	return x.transport.UpdateConstraint(ctx, v)
}

func (x *Flipt) DeleteConstraint(ctx context.Context, v *flipt.DeleteConstraintRequest) error {
	_, err := x.transport.DeleteConstraint(ctx, v)
	return err
}

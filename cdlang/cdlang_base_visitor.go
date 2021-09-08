// Code generated from CDLang.g4 by ANTLR 4.9. DO NOT EDIT.

package cdlang // CDLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCDLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCDLangVisitor) VisitContract(ctx *ContractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitConstName(ctx *ConstNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitConstProto(ctx *ConstProtoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitDomainName(ctx *DomainNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitMapping(ctx *MappingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitScenario(ctx *ScenarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitSubScenario(ctx *SubScenarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitCallResponse(ctx *CallResponseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitOpenStream(ctx *OpenStreamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitCloseStream(ctx *CloseStreamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitSend(ctx *SendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitReceive(ctx *ReceiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitExecute(ctx *ExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobuf(ctx *ProtobufContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobufField(ctx *ProtobufFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobufFieldSimple(ctx *ProtobufFieldSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitEnumer(ctx *EnumerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobufFieldGroup(ctx *ProtobufFieldGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobufFieldRepeated(ctx *ProtobufFieldRepeatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitProtobufFieldRepeatedRow(ctx *ProtobufFieldRepeatedRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCDLangVisitor) VisitPathElement(ctx *PathElementContext) interface{} {
	return v.VisitChildren(ctx)
}

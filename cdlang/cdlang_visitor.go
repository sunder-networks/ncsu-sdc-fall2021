// Code generated from CDLang.g4 by ANTLR 4.9. DO NOT EDIT.

package cdlang // CDLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CDLangParser.
type CDLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CDLangParser#contract.
	VisitContract(ctx *ContractContext) interface{}

	// Visit a parse tree produced by CDLangParser#constName.
	VisitConstName(ctx *ConstNameContext) interface{}

	// Visit a parse tree produced by CDLangParser#constProto.
	VisitConstProto(ctx *ConstProtoContext) interface{}

	// Visit a parse tree produced by CDLangParser#domainName.
	VisitDomainName(ctx *DomainNameContext) interface{}

	// Visit a parse tree produced by CDLangParser#mapping.
	VisitMapping(ctx *MappingContext) interface{}

	// Visit a parse tree produced by CDLangParser#scenario.
	VisitScenario(ctx *ScenarioContext) interface{}

	// Visit a parse tree produced by CDLangParser#subScenario.
	VisitSubScenario(ctx *SubScenarioContext) interface{}

	// Visit a parse tree produced by CDLangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by CDLangParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by CDLangParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by CDLangParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by CDLangParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by CDLangParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by CDLangParser#callResponse.
	VisitCallResponse(ctx *CallResponseContext) interface{}

	// Visit a parse tree produced by CDLangParser#openStream.
	VisitOpenStream(ctx *OpenStreamContext) interface{}

	// Visit a parse tree produced by CDLangParser#closeStream.
	VisitCloseStream(ctx *CloseStreamContext) interface{}

	// Visit a parse tree produced by CDLangParser#send.
	VisitSend(ctx *SendContext) interface{}

	// Visit a parse tree produced by CDLangParser#receive.
	VisitReceive(ctx *ReceiveContext) interface{}

	// Visit a parse tree produced by CDLangParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by CDLangParser#execute.
	VisitExecute(ctx *ExecuteContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobuf.
	VisitProtobuf(ctx *ProtobufContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobufField.
	VisitProtobufField(ctx *ProtobufFieldContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobufFieldSimple.
	VisitProtobufFieldSimple(ctx *ProtobufFieldSimpleContext) interface{}

	// Visit a parse tree produced by CDLangParser#enumer.
	VisitEnumer(ctx *EnumerContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobufFieldGroup.
	VisitProtobufFieldGroup(ctx *ProtobufFieldGroupContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobufFieldRepeated.
	VisitProtobufFieldRepeated(ctx *ProtobufFieldRepeatedContext) interface{}

	// Visit a parse tree produced by CDLangParser#protobufFieldRepeatedRow.
	VisitProtobufFieldRepeatedRow(ctx *ProtobufFieldRepeatedRowContext) interface{}

	// Visit a parse tree produced by CDLangParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by CDLangParser#pathElement.
	VisitPathElement(ctx *PathElementContext) interface{}
}

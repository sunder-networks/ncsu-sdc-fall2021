// Code generated from CDLang.g4 by ANTLR 4.9. DO NOT EDIT.

package cdlang // CDLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CDLangListener is a complete listener for a parse tree produced by CDLangParser.
type CDLangListener interface {
	antlr.ParseTreeListener

	// EnterContract is called when entering the contract production.
	EnterContract(c *ContractContext)

	// EnterConstName is called when entering the constName production.
	EnterConstName(c *ConstNameContext)

	// EnterConstProto is called when entering the constProto production.
	EnterConstProto(c *ConstProtoContext)

	// EnterDomainName is called when entering the domainName production.
	EnterDomainName(c *DomainNameContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterScenario is called when entering the scenario production.
	EnterScenario(c *ScenarioContext)

	// EnterSubScenario is called when entering the subScenario production.
	EnterSubScenario(c *SubScenarioContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterCallResponse is called when entering the callResponse production.
	EnterCallResponse(c *CallResponseContext)

	// EnterOpenStream is called when entering the openStream production.
	EnterOpenStream(c *OpenStreamContext)

	// EnterCloseStream is called when entering the closeStream production.
	EnterCloseStream(c *CloseStreamContext)

	// EnterSend is called when entering the send production.
	EnterSend(c *SendContext)

	// EnterReceive is called when entering the receive production.
	EnterReceive(c *ReceiveContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterProtobuf is called when entering the protobuf production.
	EnterProtobuf(c *ProtobufContext)

	// EnterProtobufField is called when entering the protobufField production.
	EnterProtobufField(c *ProtobufFieldContext)

	// EnterProtobufFieldSimple is called when entering the protobufFieldSimple production.
	EnterProtobufFieldSimple(c *ProtobufFieldSimpleContext)

	// EnterEnumer is called when entering the enumer production.
	EnterEnumer(c *EnumerContext)

	// EnterProtobufFieldGroup is called when entering the protobufFieldGroup production.
	EnterProtobufFieldGroup(c *ProtobufFieldGroupContext)

	// EnterProtobufFieldRepeated is called when entering the protobufFieldRepeated production.
	EnterProtobufFieldRepeated(c *ProtobufFieldRepeatedContext)

	// EnterProtobufFieldRepeatedRow is called when entering the protobufFieldRepeatedRow production.
	EnterProtobufFieldRepeatedRow(c *ProtobufFieldRepeatedRowContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterPathElement is called when entering the pathElement production.
	EnterPathElement(c *PathElementContext)

	// ExitContract is called when exiting the contract production.
	ExitContract(c *ContractContext)

	// ExitConstName is called when exiting the constName production.
	ExitConstName(c *ConstNameContext)

	// ExitConstProto is called when exiting the constProto production.
	ExitConstProto(c *ConstProtoContext)

	// ExitDomainName is called when exiting the domainName production.
	ExitDomainName(c *DomainNameContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitScenario is called when exiting the scenario production.
	ExitScenario(c *ScenarioContext)

	// ExitSubScenario is called when exiting the subScenario production.
	ExitSubScenario(c *SubScenarioContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitCallResponse is called when exiting the callResponse production.
	ExitCallResponse(c *CallResponseContext)

	// ExitOpenStream is called when exiting the openStream production.
	ExitOpenStream(c *OpenStreamContext)

	// ExitCloseStream is called when exiting the closeStream production.
	ExitCloseStream(c *CloseStreamContext)

	// ExitSend is called when exiting the send production.
	ExitSend(c *SendContext)

	// ExitReceive is called when exiting the receive production.
	ExitReceive(c *ReceiveContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitProtobuf is called when exiting the protobuf production.
	ExitProtobuf(c *ProtobufContext)

	// ExitProtobufField is called when exiting the protobufField production.
	ExitProtobufField(c *ProtobufFieldContext)

	// ExitProtobufFieldSimple is called when exiting the protobufFieldSimple production.
	ExitProtobufFieldSimple(c *ProtobufFieldSimpleContext)

	// ExitEnumer is called when exiting the enumer production.
	ExitEnumer(c *EnumerContext)

	// ExitProtobufFieldGroup is called when exiting the protobufFieldGroup production.
	ExitProtobufFieldGroup(c *ProtobufFieldGroupContext)

	// ExitProtobufFieldRepeated is called when exiting the protobufFieldRepeated production.
	ExitProtobufFieldRepeated(c *ProtobufFieldRepeatedContext)

	// ExitProtobufFieldRepeatedRow is called when exiting the protobufFieldRepeatedRow production.
	ExitProtobufFieldRepeatedRow(c *ProtobufFieldRepeatedRowContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitPathElement is called when exiting the pathElement production.
	ExitPathElement(c *PathElementContext)
}

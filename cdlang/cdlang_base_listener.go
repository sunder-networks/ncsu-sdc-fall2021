// Code generated from CDLang.g4 by ANTLR 4.9. DO NOT EDIT.

package cdlang // CDLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCDLangListener is a complete listener for a parse tree produced by CDLangParser.
type BaseCDLangListener struct{}

var _ CDLangListener = &BaseCDLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCDLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCDLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCDLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCDLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterContract is called when production contract is entered.
func (s *BaseCDLangListener) EnterContract(ctx *ContractContext) {}

// ExitContract is called when production contract is exited.
func (s *BaseCDLangListener) ExitContract(ctx *ContractContext) {}

// EnterConstName is called when production constName is entered.
func (s *BaseCDLangListener) EnterConstName(ctx *ConstNameContext) {}

// ExitConstName is called when production constName is exited.
func (s *BaseCDLangListener) ExitConstName(ctx *ConstNameContext) {}

// EnterConstProto is called when production constProto is entered.
func (s *BaseCDLangListener) EnterConstProto(ctx *ConstProtoContext) {}

// ExitConstProto is called when production constProto is exited.
func (s *BaseCDLangListener) ExitConstProto(ctx *ConstProtoContext) {}

// EnterDomainName is called when production domainName is entered.
func (s *BaseCDLangListener) EnterDomainName(ctx *DomainNameContext) {}

// ExitDomainName is called when production domainName is exited.
func (s *BaseCDLangListener) ExitDomainName(ctx *DomainNameContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BaseCDLangListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BaseCDLangListener) ExitMapping(ctx *MappingContext) {}

// EnterScenario is called when production scenario is entered.
func (s *BaseCDLangListener) EnterScenario(ctx *ScenarioContext) {}

// ExitScenario is called when production scenario is exited.
func (s *BaseCDLangListener) ExitScenario(ctx *ScenarioContext) {}

// EnterSubScenario is called when production subScenario is entered.
func (s *BaseCDLangListener) EnterSubScenario(ctx *SubScenarioContext) {}

// ExitSubScenario is called when production subScenario is exited.
func (s *BaseCDLangListener) ExitSubScenario(ctx *SubScenarioContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseCDLangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseCDLangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCDLangListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCDLangListener) ExitVariable(ctx *VariableContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCDLangListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCDLangListener) ExitConstant(ctx *ConstantContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseCDLangListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseCDLangListener) ExitVersion(ctx *VersionContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseCDLangListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseCDLangListener) ExitInstruction(ctx *InstructionContext) {}

// EnterCall is called when production call is entered.
func (s *BaseCDLangListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseCDLangListener) ExitCall(ctx *CallContext) {}

// EnterCallResponse is called when production callResponse is entered.
func (s *BaseCDLangListener) EnterCallResponse(ctx *CallResponseContext) {}

// ExitCallResponse is called when production callResponse is exited.
func (s *BaseCDLangListener) ExitCallResponse(ctx *CallResponseContext) {}

// EnterOpenStream is called when production openStream is entered.
func (s *BaseCDLangListener) EnterOpenStream(ctx *OpenStreamContext) {}

// ExitOpenStream is called when production openStream is exited.
func (s *BaseCDLangListener) ExitOpenStream(ctx *OpenStreamContext) {}

// EnterCloseStream is called when production closeStream is entered.
func (s *BaseCDLangListener) EnterCloseStream(ctx *CloseStreamContext) {}

// ExitCloseStream is called when production closeStream is exited.
func (s *BaseCDLangListener) ExitCloseStream(ctx *CloseStreamContext) {}

// EnterSend is called when production send is entered.
func (s *BaseCDLangListener) EnterSend(ctx *SendContext) {}

// ExitSend is called when production send is exited.
func (s *BaseCDLangListener) ExitSend(ctx *SendContext) {}

// EnterReceive is called when production receive is entered.
func (s *BaseCDLangListener) EnterReceive(ctx *ReceiveContext) {}

// ExitReceive is called when production receive is exited.
func (s *BaseCDLangListener) ExitReceive(ctx *ReceiveContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseCDLangListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseCDLangListener) ExitGroup(ctx *GroupContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseCDLangListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseCDLangListener) ExitExecute(ctx *ExecuteContext) {}

// EnterProtobuf is called when production protobuf is entered.
func (s *BaseCDLangListener) EnterProtobuf(ctx *ProtobufContext) {}

// ExitProtobuf is called when production protobuf is exited.
func (s *BaseCDLangListener) ExitProtobuf(ctx *ProtobufContext) {}

// EnterProtobufField is called when production protobufField is entered.
func (s *BaseCDLangListener) EnterProtobufField(ctx *ProtobufFieldContext) {}

// ExitProtobufField is called when production protobufField is exited.
func (s *BaseCDLangListener) ExitProtobufField(ctx *ProtobufFieldContext) {}

// EnterProtobufFieldSimple is called when production protobufFieldSimple is entered.
func (s *BaseCDLangListener) EnterProtobufFieldSimple(ctx *ProtobufFieldSimpleContext) {}

// ExitProtobufFieldSimple is called when production protobufFieldSimple is exited.
func (s *BaseCDLangListener) ExitProtobufFieldSimple(ctx *ProtobufFieldSimpleContext) {}

// EnterEnumer is called when production enumer is entered.
func (s *BaseCDLangListener) EnterEnumer(ctx *EnumerContext) {}

// ExitEnumer is called when production enumer is exited.
func (s *BaseCDLangListener) ExitEnumer(ctx *EnumerContext) {}

// EnterProtobufFieldGroup is called when production protobufFieldGroup is entered.
func (s *BaseCDLangListener) EnterProtobufFieldGroup(ctx *ProtobufFieldGroupContext) {}

// ExitProtobufFieldGroup is called when production protobufFieldGroup is exited.
func (s *BaseCDLangListener) ExitProtobufFieldGroup(ctx *ProtobufFieldGroupContext) {}

// EnterProtobufFieldRepeated is called when production protobufFieldRepeated is entered.
func (s *BaseCDLangListener) EnterProtobufFieldRepeated(ctx *ProtobufFieldRepeatedContext) {}

// ExitProtobufFieldRepeated is called when production protobufFieldRepeated is exited.
func (s *BaseCDLangListener) ExitProtobufFieldRepeated(ctx *ProtobufFieldRepeatedContext) {}

// EnterProtobufFieldRepeatedRow is called when production protobufFieldRepeatedRow is entered.
func (s *BaseCDLangListener) EnterProtobufFieldRepeatedRow(ctx *ProtobufFieldRepeatedRowContext) {}

// ExitProtobufFieldRepeatedRow is called when production protobufFieldRepeatedRow is exited.
func (s *BaseCDLangListener) ExitProtobufFieldRepeatedRow(ctx *ProtobufFieldRepeatedRowContext) {}

// EnterPath is called when production path is entered.
func (s *BaseCDLangListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseCDLangListener) ExitPath(ctx *PathContext) {}

// EnterPathElement is called when production pathElement is entered.
func (s *BaseCDLangListener) EnterPathElement(ctx *PathElementContext) {}

// ExitPathElement is called when production pathElement is exited.
func (s *BaseCDLangListener) ExitPathElement(ctx *PathElementContext) {}

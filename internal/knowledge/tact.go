package knowledge

import "github.com/silasdev78/silas-code-inspector/internal/domain"

func TactPatterns() []domain.Pattern {
	return []domain.Pattern{
		{ID: "TACT-001", Title: "Missing seqno check", Regex: `receive\s*\([^)]*\)\s*\{`, Severity: domain.SeverityCritical, Description: "External messages must check seqno to prevent replay attacks.", Recommendation: "Add require(self.seqno == msg.seqno) inside receive()."},
		{ID: "TACT-002", Title: "Unchecked sender", Regex: `receive\s*\(`, Severity: domain.SeverityHigh, Description: "Missing sender verification allows anyone to call this function.", Recommendation: "Add require(msg.sender == owner)."},
		{ID: "TACT-003", Title: "Unprotected selfdestruct", Regex: `selfdestruct\s*\(`, Severity: domain.SeverityCritical, Description: "selfdestruct without sender check can be called by anyone.", Recommendation: "Wrap selfdestruct in a require with owner check."},
		{ID: "TACT-004", Title: "Integer overflow/underflow", Regex: `\b(?:Int|int)\s*\.\s*(?:\+|\-|\*)`, Severity: domain.SeverityHigh, Description: "Arithmetic operations may overflow without safe math checks.", Recommendation: "Use require(result >= a) for additions or import safe math."},
		{ID: "TACT-005", Title: "Hardcoded address", Regex: `"[EQ][a-zA-Z0-9_-]{47,}"`, Severity: domain.SeverityInfo, Description: "Hardcoded address found. Ensure it's intentional.", Recommendation: "Consider using a config constant."},
		{ID: "TACT-006", Title: "Unbounded loop", Regex: `\bwhile\s*\(`, Severity: domain.SeverityHigh, Description: "Loops without gas limits can cause out-of-gas errors.", Recommendation: "Add a max iteration guard."},
		{ID: "TACT-007", Title: "External call in loop", Regex: `\.\s*send\s*\(`, Severity: domain.SeverityMedium, Description: "Sending messages inside loops consumes unpredictable gas.", Recommendation: "Batch transfers or use a single message."},
		{ID: "TACT-008", Title: "Missing storage fee consideration", Regex: `storageReserve\s*\(`, Severity: domain.SeverityMedium, Description: "Contract may lack storage fee management.", Recommendation: "Add regular storage fee top-ups."},
		{ID: "TACT-009", Title: "Timeout not enforced", Regex: `receive\s*\(`, Severity: domain.SeverityLow, Description: "No timeout logic for time-sensitive operations.", Recommendation: "Add a deadline check with now()."},
		{ID: "TACT-010", Title: "Uninitialized variable", Regex: `\bvar\s+\w+\s*:\s*\w+;`, Severity: domain.SeverityLow, Description: "Variable declared without initial value may be nil.", Recommendation: "Initialize with a sensible default."},
		{ID: "TACT-011", Title: "Gas check missing", Regex: `receive\s*\(`, Severity: domain.SeverityMedium, Description: "No gas-usage limit set for receive function.", Recommendation: "Add require(gasConsumed < MAX_GAS)."},
		{ID: "TACT-012", Title: "Old compiler version", Regex: `^\s*\/\/\s*compiler\s*:`, Severity: domain.SeverityLow, Description: "Compiler version might be outdated, leading to bugs.", Recommendation: "Update to latest Tact compiler."},
		{ID: "TACT-013", Title: "Signature verification missing", Regex: `receive\s*\(`, Severity: domain.SeverityHigh, Description: "No signature verification for critical actions.", Recommendation: "Use checkDataSign or similar."},
		{ID: "TACT-014", Title: "Unchecked return value", Regex: `\.\s*send\s*\(`, Severity: domain.SeverityMedium, Description: "send() return flag ignored; message might fail silently.", Recommendation: "Check the boolean return of send()."},
		{ID: "TACT-015", Title: "Missing state isolation", Regex: `\bstate\b\s*\{`, Severity: domain.SeverityLow, Description: "State declaration without clear isolation comment.", Recommendation: "Document state invariants."},
		{ID: "TACT-016", Title: "Insecure random source", Regex: `\brandom\s*\(`, Severity: domain.SeverityHigh, Description: "Randomness may be predictable from block data.", Recommendation: "Use commitment schemes or off-chain randomness."},
		{ID: "TACT-017", Title: "Lack of access control modifier", Regex: `\bfun\s+\w+\s*\(`, Severity: domain.SeverityMedium, Description: "Functions without access modifiers may be public by default.", Recommendation: "Explicitly declare visibility."},
		{ID: "TACT-018", Title: "Exposed internal function", Regex: `\bpublic\s+fun\s+internal\w*`, Severity: domain.SeverityInfo, Description: "Naming suggests internal logic but marked public.", Recommendation: "Make private or clearly document."},
		{ID: "TACT-019", Title: "Unchecked msg value", Regex: `receive\s*\(`, Severity: domain.SeverityLow, Description: "No check on msg.value allows zero-value calls.", Recommendation: "Add require(msg.value > 0)."},
		{ID: "TACT-020", Title: "Missing code upgrade path", Regex: `\bcode\s*:`, Severity: domain.SeverityInfo, Description: "No mechanism for code upgrade, contract is immutable.", Recommendation: "Consider a proxy pattern or upgradeable state."},
	}
}

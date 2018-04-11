package stc

// DebugPackagePng is the base64 encoding of the
// debug package png icon
const DebugPackagePng = `
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
`

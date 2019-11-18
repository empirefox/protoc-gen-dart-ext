package dartpb

func (l *L10nMsgOrEnum) Ignore() bool {
	return l.Extension.Ignore
}
func (l *L10nOneOf) Ignore() bool {
	return l.Entity.Message.L10n.Ignore() || l.Extension.Ignore
}
func (l *L10nEnumValue) Ignore() bool {
	return l.Entity.Enum.L10n.Ignore() || l.Extension.Ignore
}
func (l *L10nField) Ignore() bool {
	return l.Entity.Message.L10n.Ignore() || l.Extension.Ignore
}

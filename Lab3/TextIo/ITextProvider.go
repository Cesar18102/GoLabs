package TextIo

type ITextProvider interface {
	Write(text string);
	Read() string;
}
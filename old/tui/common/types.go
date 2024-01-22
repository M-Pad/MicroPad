package common

type Intractable interface {
	Focus()
	Blur()
}

var NormalChars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!\"#$%&\\'()*+,-./:;<=>?@[\\]^_`{|}~ "

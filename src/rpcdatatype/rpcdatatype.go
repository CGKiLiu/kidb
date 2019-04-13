package rpcdatatype


type Load struct{
	Key []byte
}

type Store struct{
	Key  []byte
	Value []byte
}

type NullResult int

type ValueResult struct{
	Value []byte
}
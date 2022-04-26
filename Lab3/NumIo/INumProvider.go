package NumIo

type INumProvider interface {
	Add(n int);
	RemoveAll(n int);
	RemoveFirst(n int) bool;
	Clear();
	Commit();
	Get() []int;
}

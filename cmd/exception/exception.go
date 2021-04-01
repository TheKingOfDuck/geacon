package exception

//Exception exception type include err, panic info, and other...
type Exception interface{}

//Try things try to do
type Try func()

//Catch deal with exception
type Catch func(e Exception)

//Trier a obj like try...catch...
type Trier struct {
	exception Exception
}

//New return trier
func New() *Trier {
	return &Trier{}
}

//Try to do things
func (t *Trier) Try(f Try) (trier *Trier) {
	trier = t
	defer func() {
		if err := recover(); err != nil {
			t.exception = err
		}
	}()

	f()
	return t
}

//Throw a exception
func (t *Trier) Throw(e Exception) {
	if e != nil {
		panic(e)
	}
}

//Catch exception and deal with it
func (t *Trier) Catch(f Catch) {
	if t.exception != nil {
		f(t.exception)
	}
}
package collection

type hashSet[E comparable] struct {
	m Map[E, struct{}]
}

func NewHashSet[E comparable]() Set[E] {
	m := NewHashMap[E, struct{}]()
	return &hashSet[E]{
		m: m,
	}
}

func NewHashSetFrom[E comparable](c Collection[E]) Set[E] {
	m := NewHashMap[E, struct{}]()
	c.Range(func(e E) bool {
		m.Store(e, struct{}{})
		return true
	})
	return &hashSet[E]{
		m: m,
	}
}

func NewSyncHashSet[E comparable]() Set[E] {
	m := NewSyncHashMap[E, struct{}]()
	return &hashSet[E]{
		m: m,
	}
}

func NewSyncHashSetFrom[E comparable](c Collection[E]) Set[E] {
	m := NewSyncHashMap[E, struct{}]()
	c.Range(func(e E) bool {
		m.Store(e, struct{}{})
		return true
	})
	return &hashSet[E]{
		m: m,
	}
}

func (s *hashSet[E]) Add(e ...E) bool {
	changed := false
	for _, ee := range e {
		_, ok := s.m.LoadAndStore(ee, struct{}{})
		if ok {
			changed = true
		}
	}
	return changed
}

func (s *hashSet[E]) AddAll(c Collection[E]) bool {
	changed := false
	c.Range(func(e E) bool {
		_, ok := s.m.LoadAndStore(e, struct{}{})
		if ok {
			changed = true
		}
		return true
	})
	return changed
}

func (s *hashSet[E]) Contains(e E) bool {
	_, ok := s.m.Load(e)
	return ok
}

func (s *hashSet[E]) Size() int {
	return s.m.Size()
}

func (s *hashSet[E]) IsEmpty() bool {
	return s.m.IsEmpty()
}

func (s *hashSet[E]) Delete(e E) bool {
	_, ok := s.m.LoadAndDelete(e)
	return ok
}

func (s *hashSet[E]) Clear() {
	s.m.Clear()
}

func (s *hashSet[E]) Range(f func(e E) bool) {
	s.m.Range(func(key E, value struct{}) bool {
		return f(key)
	})
}

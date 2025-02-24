package todo

type mockStorage struct {
	Storager
	id   int64
	todo Todo
	err  error
}

func (m mockStorage) Todos() ([]Todo, error) {
	return []Todo{m.todo}, m.err
}

func (m *mockStorage) Todo(id int64) (Todo, error) {
	m.id = id
	return m.todo, m.err
}

func (m *mockStorage) Create(todo Todo) (int64, error) {
	m.todo = todo
	return m.id, m.err
}

func (m *mockStorage) Update(todo Todo) error {
	m.todo = todo
	return m.err
}

func (m *mockStorage) Delete(id int64) error {
	m.id = id
	return m.err
}

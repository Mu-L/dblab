package pagination

import "fmt"

// Manager handles the pagination.
type Manager struct {
	totalPages  int
	currentPage int
	limit       int
	offset      int
	active      bool
}

// New returns a pointer to a Manager instance.
func New(limit, count int) (*Manager, error) {
	m := Manager{
		limit:       limit,
		currentPage: 1,
		active:      true,
	}

	m.setOffset()

	if err := m.setTotalPages(count); err != nil {
		return nil, err
	}

	return &m, nil
}

// NextPage increases the value currentPage.
func (m *Manager) NextPage() error {
	if m.currentPage+1 > m.totalPages || !m.active {
		return fmt.Errorf("current page should not be greater than the total pages count")
	}

	m.currentPage++
	m.setOffset()

	return nil
}

// PreviousPage decreases the value of currentPage.
func (m *Manager) PreviousPage() error {
	if m.currentPage-1 <= 0 || !m.active {
		return fmt.Errorf("current page should not be less than 0")
	}

	m.currentPage--
	m.setOffset()

	return nil
}

// Deactivate sets the active property to false.
func (m *Manager) Deactivate() {
	m.active = false
}

// Offset returns the limit.
func (m *Manager) Offset() int {
	return m.offset
}

// Limit returns the limit.
func (m *Manager) Limit() int {
	return m.limit
}

// setOffset calculates the offset based of the current page and the limit.
func (m *Manager) setOffset() {
	m.offset = (m.currentPage - 1) * m.limit
}

// CurrentPage returns the currentPage value.
func (m *Manager) CurrentPage() int {
	return m.currentPage
}

// setTotalPages total pages = count / limit, if the limit is greater than 0.
func (m *Manager) setTotalPages(count int) error {
	// limit must be greater than 0.
	if m.limit <= 0 {
		return fmt.Errorf("limit should be greater than 0")
	}

	m.totalPages = count / m.limit

	return nil
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Interfaces to maintain rows

// Row interface defines the implementation needed for SerializeKey & SerializeRow (The implementation would differ for each Table type).
type Row interface {
	SerializeRow() []byte
	SerializeKey() []byte
	GetKey() interface{}
}

// Table schemas and implements Row interface
type smCtxRow struct {
	SmCtxRef string
	Supi     string
	PduSesID string
}

func (row smCtxRow) SerializeRow() []byte {
	// Implement the serialization logic for Row in smCtxRefTable
	return nil
}

func (row smCtxRow) SerializeKey() []byte {
	// Implement the serialization logic for key in smCtxRefTable
	return nil
}

func (row smCtxRow) GetKey() interface{} {
	// Implement the serialization logic for key in smCtxRefTable
	return nil
}

//Note: Similarly PeiToSupiTable and StatusNotifyToSupiTable must be implement - Row interface.
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type RowDeserializer interface {
	DeserializeRow([]byte) Row
}

// implements RowDeserializer interface
type smCtxRefTable struct {
}

func (table smCtxRefTable) DeserializeRow(data []byte) Row {
	var recoveredRow smCtxRow
	json.Unmarshal(data, &recoveredRow)
	return recoveredRow
}

type IMirrorDB interface {
	DeserializeRow([]byte) Row
}

type Table struct { // implements IMirrorDB interface
	rows         map[interface{}]Row
	deserialImpl RowDeserializer
}

func (t *Table) DeserializeRow(serializedData []byte) Row {
	return t.deserialImpl.DeserializeRow(serializedData)
}

//Note: Similarly RowDeserializer must be implemented for PeiToSupiTable and StatusNotifyToSupiTable
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Map storing the reference of each table
var tables map[string]*Table

func (t *Table) AddRow(key interface{}, row Row) {
	key = row.GetKey()
	t.rows[key] = row
}

func (t *Table) GetRow(key interface{}) Row {
	return t.rows[key]
}

func (t *Table) DeleteRow(key interface{}) {
	delete(t.rows, key)
}

func (t *Table) CreateTable(tableName string, rowDeserializer RowDeserializer) *Table {
	m := make(map[interface{}]Row)
	table := Table{rows: m, deserialImpl: rowDeserializer}
	tables[tableName] = &table
	return &table
}

func (t *Table) DeleteTable(tableName string) {
	delete(tables, tableName)
}

func (t *Table) GetTable(tableName string) *Table {
	return tables[tableName]
}

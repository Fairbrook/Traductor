package Semantico

import "github.com/Fairbrook/analizador/Utils"

type TableEntry struct {
	symbol Symbol
	table  *Table
}

type Table struct {
	ts     map[string]TableEntry
	Stack  Utils.Stack
	Parent Symbol
}

func (t *Table) Set(symbol Symbol, subtable *Table) {
	if t.ts == nil {
		t.ts = map[string]TableEntry{}
	}
	t.ts[symbol.GetIdentifier()] = TableEntry{symbol: symbol, table: subtable}
	t.Stack.Push(symbol)
}

func (t *Table) Get(identifier string) (Symbol, bool) {
	item, ok := t.ts[identifier]
	if ok {
		return item.symbol, true
	}
	if t.Parent != nil && t.Parent.GetIdentifier() == identifier {
		return t.Parent, false
	}
	iterator := t.Stack.GetListPointer()
	for iterator != nil {
		if iterator.Data.(Symbol).GetIdentifier() == identifier {
			return iterator.Data.(Symbol), false
		}
		iterator = iterator.Next
	}
	return nil, false
}

func (t *Table) Includes(identifier string, sameScope bool) bool {
	_, ok := t.ts[identifier]
	if ok {
		return true
	}
	if t.Parent != nil && t.Parent.GetIdentifier() == identifier {
		return true
	}
	if sameScope {
		return false
	}
	iterator := t.Stack.GetListPointer()
	for iterator != nil {
		if iterator.Data.(Symbol).GetIdentifier() == identifier {
			return true
		}
		iterator = iterator.Next
	}
	return false
}

func (t *Table) dumpStack() Utils.Stack {
	copy := Utils.Stack{}
	iterator := t.Stack.GetListPointer()
	for iterator != nil {
		copy.Push(iterator.Data)
		iterator = iterator.Next
	}
	return copy
}

func (t *Table) ToArray() [][3]string {
	return t.toArrayInter("")
}

func (t *Table) toArrayInter(parent string) [][3]string {
	prepend := ""
	res := [][3]string{}
	if parent != "" {
		prepend = parent + " ➔ "
	}
	for key, te := range t.ts {
		temp := te.symbol.ToArray()
		temp[0] = prepend + temp[0]
		res = append(res, temp)
		if te.table != nil {
			res = append(res, te.table.toArrayInter(prepend+key)...)
		}
	}
	return res
}

func (t *Table) DumpTable() []Symbol {
	res := []Symbol{}
	for _, val := range t.ts {
		res = append(res, val.symbol)
	}
	return res
}

func (t *Table) GetSubTable(identifier string) *Table {
	item, ok := t.ts[identifier]
	if ok {
		return item.table
	}
	return nil
}

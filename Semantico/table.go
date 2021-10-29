package Semantico

import "github.com/Fairbrook/analizador/Utils"

type TableEntry struct {
	symbol Symbol
	table  *Table
}

type Table struct {
	ts    map[string]TableEntry
	Stack Utils.Stack
}

func (t *Table) Set(symbol Symbol, subtable *Table) {
	if t.ts == nil {
		t.ts = map[string]TableEntry{}
	}
	t.ts[symbol.getIdentifier()] = TableEntry{symbol: symbol, table: subtable}
	t.Stack.Push(symbol)
}

func (t *Table) Get(identifier string) Symbol {
	item, ok := t.ts[identifier]
	if ok {
		return item.symbol
	}
	iterator := t.Stack.GetListPointer()
	for iterator != nil {
		if iterator.Data.(Symbol).getIdentifier() == identifier {
			return iterator.Data.(Symbol)
		}
		iterator = iterator.Next
	}
	return nil
}

func (t *Table) Includes(identifier string) bool {
	_, ok := t.ts[identifier]
	if ok {
		return true
	}
	iterator := t.Stack.GetListPointer()
	for iterator != nil {
		if iterator.Data.(Symbol).getIdentifier() == identifier {
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

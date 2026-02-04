package main

import (
	"database/sql"
	"fmt"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

// Add создаёт новую посылку в БД и возвращает её номер
func (s ParcelStore) Add(p Parcel) (int, error) {
	res, err := s.db.Exec("INSERT INTO parcel (client, status, address, created_at) VALUES (?, ?, ?, ?)",
		p.Client,
		p.Status,
		p.Address,
		p.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get возвращает посылку по её номеру
func (s ParcelStore) Get(number int) (Parcel, error) {
	p := Parcel{}
	row := s.db.QueryRow("SELECT number, client, status, address, created_at FROM parcel WHERE number = ?", number)
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		return Parcel{}, err
	}
	return p, nil
}

// GetByClient возвращает все посылки указанного клиента
func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	rows, err := s.db.Query("SELECT number, client, status, address, created_at FROM parcel WHERE client = ?", client)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Parcel
	for rows.Next() {
		p := Parcel{}
		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// SetStatus обновляет статус посылки по номеру
func (s ParcelStore) SetStatus(number int, status string) error {
	_, err := s.db.Exec("UPDATE parcel SET status = ? WHERE number = ?",
		status,
		number,
	)
	if err != nil {
		return err
	}
	return nil
}

// SetAddress изменяет адрес доставки, если посылка находится в статусе "registered"
func (s ParcelStore) SetAddress(number int, address string) error {
	parcel, err := s.Get(number)
	if err != nil {
		return err
	}
	if parcel.Status != ParcelStatusRegistered {
		return fmt.Errorf("can't change address because parcel status: %s", parcel.Status)
	}
	_, err = s.db.Exec("UPDATE parcel SET address = ? WHERE number = ?",
		address,
		number,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete удаляет посылку, если она находится в статусе "registered"
func (s ParcelStore) Delete(number int) error {
	parcel, err := s.Get(number)
	if err != nil {
		return err
	}
	if parcel.Status != ParcelStatusRegistered {
		return fmt.Errorf("can't delete row because parcel status %s", parcel.Status)
	}
	_, err = s.db.Exec("DELETE FROM parcel WHERE number = ?", number)
	if err != nil {
		return err
	}
	return nil
}

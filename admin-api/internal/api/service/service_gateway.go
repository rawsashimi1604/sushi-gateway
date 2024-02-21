package service

import "github.com/jackc/pgx/v5/pgxpool"

type ServiceGateway struct {
	Conn *pgxpool.Pool
}

func NewServiceGateway(conn *pgxpool.Pool) *ServiceGateway {
	return &ServiceGateway{
		Conn: conn,
	}
}

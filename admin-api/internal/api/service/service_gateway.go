package service

//
//import (
//	"context"
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//type ServiceGateway struct {
//	Conn *pgxpool.Pool
//}
//
//func NewServiceGateway(conn *pgxpool.Pool) *ServiceGateway {
//	return &ServiceGateway{
//		Conn: conn,
//	}
//}
//
//func (sg *ServiceGateway) getAllServices() ([]Service, error) {
//	query := `
//		SELECT s.id, s.created_at, s.updated_at, s.name, s.host, s.port,
//		       s.protocol, s.tags, s.enabled, s.health_check_enabled, s.health
//		FROM services s
//-- 		LEFT JOIN routes r
//-- 		ON s.id=r.service_id
//	`
//
//	// Query the database
//	rows, err := sg.Conn.Query(context.Background(), query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	// Iterate through the rows
//	for rows.Next() {
//		var s Service
//		if err := rows.Scan(&s.Id, &s.CreatedAt, &s.UpdatedAt, &s.); err != nil {
//			return nil, err
//		}
//		products = append(products, p)
//	}
//
//	// Check for any errors encountered during iteration
//	if rows.Err() != nil {
//		return nil, err
//	}
//
//	return products, nil
//
//	return make([]Service, 0), nil
//}

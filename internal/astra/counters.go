package astra

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
)

func (c *Client) CreateCountersTable() error {
	query := &pb.Query{
		Cql: fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
			id UUID PRIMARY KEY,
			count counter
			);`, c.keyspace, c.countersTable),
	}
	_, err := c.sgClient.ExecuteQuery(query)
	return err
}

func (c *Client) Increment(id string, inc int) error {

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	query := &pb.Query{
		Cql: fmt.Sprintf(`UPDATE %s.%s SET count = count + %d WHERE id = %s;`,
			c.keyspace, c.countersTable, inc, uid.String()),
	}
	_, err = c.sgClient.ExecuteQuery(query)
	return err
}

func (c *Client) Lookup(id string) (int, error) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return 0, err
	}

	query := &pb.Query{
		Cql: fmt.Sprintf(
			`SELECT count FROM %s.%s WHERE id = %s LIMIT 1;`,
			c.keyspace, c.countersTable, uid.String()),
	}

	response, err := c.sgClient.ExecuteQuery(query)
	if err != nil {
		return 0, err
	}

	result := response.GetResultSet()

	if len(result.Rows) < 1 {
		return 0, errors.New("id not found")
	}

	countVal := result.Rows[0].Values[0]
	return int(countVal.GetInt()), nil
}

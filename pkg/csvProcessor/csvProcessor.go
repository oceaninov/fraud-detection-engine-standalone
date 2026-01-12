package csvProcessor

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

type DTTOTCSVFormat struct {
	PPATKID     string
	Name        string
	NIK         string
	DateOfBirth string
	DataSource  string
}

type MerchantCSVFormat struct {
	NMID       string
	Name       string
	DataSource string
}

type TransactionCSVFormat struct {
	Id                  string
	TransactionId       string
	TransactionType     string
	Title               string
	Channel             string
	Flag                string
	UserId              string
	Amount              string
	DestinationId       string
	CreatedAt           string
	RuleName            string
	RuleAmount          string
	RuleInterval        string
	RuleTransactionType string
	RuleAction          string
	BodyReq             string
}

type BlacklistCSVFormat struct {
	Id               string
	PhoneNumber      string
	BeneficiaryName  string
	TransactionTypes string
	Event            string
	CreatedAt        string
	CreatedBy        string
}

func ReadCSVForBlacklist(bytesData []byte) ([]BlacklistCSVFormat, error) {
	// Create a CSV reader from the byte slice
	reader := csv.NewReader(strings.NewReader(string(bytesData)))
	reader.Comma = ';' // Set the delimiter to semicolon (;)

	// Read all rows from the CSV data
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %v", err)
	}

	// Parse rows into a slice of DTTOTCSVFormat
	var records []BlacklistCSVFormat
	for i, row := range rows {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) != 6 {
			return nil, fmt.Errorf("unexpected column count at row %d", i+1)
		}
		records = append(records, BlacklistCSVFormat{
			Id:               row[0],
			PhoneNumber:      row[1],
			BeneficiaryName:  row[2],
			TransactionTypes: row[3],
			Event:            row[4],
			CreatedAt:        row[5],
			CreatedBy:        row[6],
		})
	}

	return records, nil
}

func ReadCSVForTransaction(bytesData []byte) ([]TransactionCSVFormat, error) {
	// Create a CSV reader from the byte slice
	reader := csv.NewReader(strings.NewReader(string(bytesData)))
	reader.Comma = ';' // Set the delimiter to semicolon (;)

	// Read all rows from the CSV data
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %v", err)
	}

	// Parse rows into a slice of DTTOTCSVFormat
	var records []TransactionCSVFormat
	for i, row := range rows {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) != 15 {
			return nil, fmt.Errorf("unexpected column count at row %d", i+1)
		}
		records = append(records, TransactionCSVFormat{
			Id:                  row[0],
			TransactionId:       row[1],
			TransactionType:     row[2],
			Title:               row[3],
			Channel:             row[4],
			Flag:                row[5],
			UserId:              row[6],
			Amount:              row[7],
			DestinationId:       row[8],
			CreatedAt:           row[9],
			RuleName:            row[10],
			RuleAmount:          row[11],
			RuleInterval:        row[12],
			RuleTransactionType: row[13],
			RuleAction:          row[14],
			BodyReq:             row[15],
		})
	}

	return records, nil
}

func ReadCSVForDTTOT(bytesData []byte) ([]DTTOTCSVFormat, error) {
	// Create a CSV reader from the byte slice
	reader := csv.NewReader(strings.NewReader(string(bytesData)))
	reader.Comma = ';' // Set the delimiter to semicolon (;)

	// Read all rows from the CSV data
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %v", err)
	}

	// Parse rows into a slice of DTTOTCSVFormat
	var records []DTTOTCSVFormat
	for i, row := range rows {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) != 5 {
			return nil, fmt.Errorf("unexpected column count at row %d", i+1)
		}
		records = append(records, DTTOTCSVFormat{
			PPATKID:     row[0],
			Name:        row[1],
			NIK:         row[2],
			DateOfBirth: row[3],
			DataSource:  row[4],
		})
	}

	return records, nil
}

func ReadCSVForMerchant(bytesData []byte) ([]MerchantCSVFormat, error) {
	// Create a CSV reader from the byte slice
	reader := csv.NewReader(strings.NewReader(string(bytesData)))
	reader.Comma = ';' // Set the delimiter to semicolon (;)

	// Read all rows from the CSV data
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %v", err)
	}

	// Parse rows into a slice of DTTOTCSVFormat
	var records []MerchantCSVFormat
	for i, row := range rows {
		if i == 0 {
			// Skip the header row
			continue
		}
		if len(row) != 3 {
			return nil, fmt.Errorf("unexpected column count at row %d", i+1)
		}
		records = append(records, MerchantCSVFormat{
			NMID:       row[0],
			Name:       row[1],
			DataSource: row[2],
		})
	}

	return records, nil
}

func ConvertTransactionToCSV(records []TransactionCSVFormat) (string, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	writer.Comma = ';' // Set the delimiter to semicolon (;)

	// Write header
	header := []string{
		"Id",
		"TransactionId",
		"TransactionType",
		"Title",
		"Channel",
		"Flag",
		"UserId",
		"Amount",
		"DestinationId",
		"CreatedAt",
		"RuleName",
		"RuleAmount",
		"RuleInterval",
		"RuleTransactionType",
		"RuleAction",
		"BodyRe",
	}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write header: %v", err)
	}

	// Write records
	for _, record := range records {
		row := []string{
			record.Id,
			record.TransactionId,
			record.TransactionType,
			record.Title,
			record.Channel,
			record.Flag,
			record.UserId,
			record.Amount,
			record.DestinationId,
			record.CreatedAt,
			record.RuleName,
			record.RuleAmount,
			record.RuleInterval,
			record.RuleTransactionType,
			record.RuleAction,
			record.BodyReq,
		}
		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write record: %v", err)
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("error while flushing writer: %v", err)
	}

	return buffer.String(), nil
}

func ConvertBlacklistToCSV(records []BlacklistCSVFormat) (string, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	writer.Comma = ';' // Set the delimiter to semicolon (;)

	// Write header
	header := []string{
		"Id",
		"PhoneNumber",
		"BeneficiaryName",
		"TransactionTypes",
		"Event",
		"CreatedAt",
		"CreatedBy",
	}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write header: %v", err)
	}

	// Write records
	for _, record := range records {
		row := []string{
			record.Id,
			record.PhoneNumber,
			record.BeneficiaryName,
			record.TransactionTypes,
			record.Event,
			record.CreatedAt,
			record.CreatedBy,
		}
		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write record: %v", err)
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("error while flushing writer: %v", err)
	}

	return buffer.String(), nil
}

func ConvertDTTOTToCSV(records []DTTOTCSVFormat) (string, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	writer.Comma = ';' // Set the delimiter to semicolon (;)

	// Write header
	header := []string{"PPATKID", "Name", "NIK", "DateOfBirth", "DataSource"}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write header: %v", err)
	}

	// Write records
	for _, record := range records {
		row := []string{record.PPATKID, record.Name, record.NIK, record.DateOfBirth, record.DataSource}
		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write record: %v", err)
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("error while flushing writer: %v", err)
	}

	return buffer.String(), nil
}

func ConvertMerchantToCSV(records []MerchantCSVFormat) (string, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	writer.Comma = ';' // Set the delimiter to semicolon (;)

	// Write header
	header := []string{"NMID", "Name", "DataSource"}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write header: %v", err)
	}

	// Write records
	for _, record := range records {
		row := []string{record.NMID, record.Name, record.DataSource}
		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write record: %v", err)
		}
	}

	// Flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("error while flushing writer: %v", err)
	}

	return buffer.String(), nil
}

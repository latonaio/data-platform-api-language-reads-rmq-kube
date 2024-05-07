package dpfm_api_output_formatter

import (
	"data-platform-api-language-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToLanguage(rows *sql.Rows) (*[]Language, error) {
	defer rows.Close()
	language := make([]Language, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Language{}

		err := rows.Scan(
			&pm.Language,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &language, nil
		}

		data := pm
		language = append(language, Language{
			Language:            data.Language,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &language, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.Language,
			&pm.CorrespondenceLanguage,
			&pm.LanguageName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			Language:               data.Language,
			CorrespondenceLanguage: data.CorrespondenceLanguage,
			LanguageName:           data.LanguageName,
			CreationDate:           data.CreationDate,
			LastChangeDate:         data.LastChangeDate,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}

package models

func ConvertToGroup(resultGroup []ResultGroup) []Group {
	var groups []Group

	for _, rg := range resultGroup {
		group := Group{
			Id:           rg.Id,
			Image:        rg.Image,
			Name:         rg.Name,
			Members:      rg.Members,
			CreationDate: rg.CreationDate,
			FirstAlbum:   rg.FirstAlbum,
		}

		groups = append(groups, group)
	}

	return groups
}

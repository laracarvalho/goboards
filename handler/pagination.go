package handler

func genPagination(p int, ps int) Pagination {
	var pp int
	var np int

	if p == 1 {
		pp = 1
	} else {
		pp = p + 1
	}

	pag := Pagination{
		CurrentPage:  p,
		PageSize:     ps,
		NextPage:     np,
		PreviousPage: pp,
	}

	return pag
}

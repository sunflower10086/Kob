package all

import (
	_ "backend/internal/bot/http"
	_ "backend/internal/bot/impl"

	_ "backend/internal/user/http"
	_ "backend/internal/user/impl"

	_ "backend/internal/record/http"
	_ "backend/internal/record/impl"

	_ "backend/internal/rank/http"
	_ "backend/internal/rank/impl"
)

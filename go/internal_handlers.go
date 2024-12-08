package main

import (
	"database/sql"
	"errors"
	"math/rand"
	"net/http"
)

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := db.Beginx()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	defer tx.Rollback()

	ride := &Ride{}
	if err := tx.GetContext(ctx, ride, `SELECT * FROM rides WHERE chair_id IS NULL ORDER BY created_at LIMIT 1 FOR UPDATE`); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	var matched *Chair
	found := false

	for i := 0; i < 10; i++ {
		var total int
		// 椅子の総数を取得
		if err := tx.GetContext(ctx, &total, "SELECT COUNT(*) FROM chairs WHERE is_active = TRUE"); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		if total == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		randOffset := rand.Intn(total)
		matched = &Chair{}
		// ランダムな椅子を取得
		if err := tx.GetContext(ctx, matched, "SELECT * FROM chairs WHERE is_active = TRUE LIMIT 1 OFFSET ?", randOffset); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		// 椅子をライドに割り当てる
		if _, err := tx.ExecContext(ctx, "UPDATE rides SET chair_id = ? WHERE id = ?", matched.ID, ride.ID); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		found = true
		break
	}

	if !found {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := tx.Commit(); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

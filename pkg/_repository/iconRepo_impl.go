package _repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
	"github.com/DueIt-Jasanya-Aturuang/toy-story/util"
)

type IconRepositoryImpl struct {
	db *sql.DB
}

func NewIconRepositoryImpl(db *sql.DB) domain.IconRepository {
	return &IconRepositoryImpl{
		db: db,
	}
}

func (i *IconRepositoryImpl) Create(ctx context.Context, payment *domain.Icon) error {
	query := `INSERT INTO m_icon (id, title, icon, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		payment.ID,
		payment.Title,
		payment.Icon,
		payment.CreatedAt,
		payment.CreatedBy,
		payment.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IconRepositoryImpl) Update(ctx context.Context, payment *domain.Icon) error {
	query := `UPDATE m_icon SET title = $1, icon = $2, updated_at = $3, updated_by = $4
				WHERE id = $5`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		payment.Title,
		payment.Icon,
		payment.UpdatedAt,
		payment.UpdatedBy,
		payment.ID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IconRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `UPDATE m_icon SET deleted_at = $1, deleted_by = $2
				WHERE id = $3 AND deleted_at IS NULL`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		time.Now().Unix(),
		"",
		id,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IconRepositoryImpl) GetAll(ctx context.Context) (*[]domain.Icon, error) {
	query := `SELECT id, title, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
				WHERE deleted_at IS NULL`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return nil, err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, errClose)
		}
	}()

	var icons []domain.Icon
	var icon domain.Icon

	for rows.Next() {
		if err = rows.Scan(
			&icon.ID,
			&icon.Title,
			&icon.Icon,
			&icon.CreatedAt,
			&icon.CreatedBy,
			&icon.UpdatedAt,
			&icon.UpdatedBy,
			&icon.DeletedAt,
			&icon.DeletedBy,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}

		icons = append(icons, icon)
	}

	return &icons, nil
}

func (i *IconRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Icon, error) {
	query := `SELECT id, title, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
				WHERE id = $1 AND deleted_at IS NULL`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return nil, err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var icon domain.Icon
	if err = stmt.QueryRowContext(ctx, id).Scan(
		&icon.ID,
		&icon.Title,
		&icon.Icon,
		&icon.CreatedAt,
		&icon.CreatedBy,
		&icon.UpdatedAt,
		&icon.UpdatedBy,
		&icon.DeletedAt,
		&icon.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
		}
		return nil, err
	}

	return &icon, nil
}

func (i *IconRepositoryImpl) GetByTitle(ctx context.Context, title string) (*domain.Icon, error) {
	query := `SELECT id, title, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
				WHERE title = $1 AND deleted_at IS NULL`

	conn, err := i.db.Conn(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrDBConn, err)
		return nil, err
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrDBConnClose, errClose)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var icon domain.Icon
	if err = stmt.QueryRowContext(ctx, title).Scan(
		&icon.ID,
		&icon.Title,
		&icon.Icon,
		&icon.CreatedAt,
		&icon.CreatedBy,
		&icon.UpdatedAt,
		&icon.UpdatedBy,
		&icon.DeletedAt,
		&icon.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
		}
		return nil, err
	}

	return &icon, nil
}

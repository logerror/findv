package db

import (
	bolt "go.etcd.io/bbolt"
	"golang.org/x/xerrors"
	"os"
	"path/filepath"
	"runtime/debug"
)

var (
	db    *bolt.DB
	dbDir string
)

func InitDbWithPath(cacheDir string) (err error) {
	dbPath := Path(cacheDir)
	dbDir = filepath.Dir(dbPath)
	if err = os.MkdirAll(dbDir, 0700); err != nil {
		return xerrors.Errorf("failed to mkdir: %w", err)
	}

	// bbolt sometimes occurs the fatal error of "unexpected fault address".
	// In that case, the local DB should be broken and needs to be removed.
	debug.SetPanicOnFault(true)
	defer func() {
		if r := recover(); r != nil {
			if err = os.Remove(dbPath); err != nil {
				return
			}
			db, err = bolt.Open(dbPath, 0600, nil)
		}
		debug.SetPanicOnFault(false)
	}()

	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return xerrors.Errorf("failed to open db: %w", err)
	}
	return nil
}

func Dir(cacheDir string) string {
	return filepath.Join(cacheDir, "db")
}

func Path(cacheDir string) string {
	dbPath := filepath.Join(Dir(cacheDir), "trivy.db")
	return dbPath
}

func Close() error {
	// Skip closing the database if the connection is not established.
	if db == nil {
		return nil
	}
	if err := db.Close(); err != nil {
		return xerrors.Errorf("failed to close DB: %w", err)
	}
	return nil
}

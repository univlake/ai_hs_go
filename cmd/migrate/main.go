package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dsn := flag.String("dsn", os.Getenv("PG_DSN"), "PostgreSQL DSN (postgres://...)")
	dir := flag.String("dir", "./migrations", "migrations directory")
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}
	cmd := flag.Arg(0)

	switch cmd {
	case "create":
		if err := createCmd(*dir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "up", "down", "version", "force":
		if *dsn == "" {
			fmt.Fprintln(os.Stderr, "missing -dsn (or PG_DSN env)")
			os.Exit(1)
		}
		if err := runMigrate(cmd, *dsn, *dir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: migrate [-dsn ...] [-dir ./migrations] {up|down|version|force|create}")
	fmt.Fprintln(os.Stderr, "  create -name <name>   write empty up/down sql files")
	fmt.Fprintln(os.Stderr, "  up                    apply all pending migrations")
	fmt.Fprintln(os.Stderr, "  down -n N             roll back N migrations")
	fmt.Fprintln(os.Stderr, "  version               print current migration version")
	fmt.Fprintln(os.Stderr, "  force -v V            force version V (recover from dirty state)")
}

func runMigrate(cmd, dsn, dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("abs dir: %w", err)
	}
	src := "file://" + abs
	m, err := migrate.New(src, dsn)
	if err != nil {
		return fmt.Errorf("init: %w", err)
	}
	defer m.Close()

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		fmt.Println("migrate up: ok")
	case "down":
		fs := flag.NewFlagSet("down", flag.ExitOnError)
		n := fs.Int("n", 1, "number of migrations to roll back")
		_ = fs.Parse(flag.Args()[1:])
		if err := m.Steps(-*n); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		fmt.Printf("migrate down -%d: ok\n", *n)
	case "version":
		v, dirty, err := m.Version()
		if err != nil {
			return err
		}
		fmt.Printf("version=%d dirty=%v\n", v, dirty)
	case "force":
		fs := flag.NewFlagSet("force", flag.ExitOnError)
		v := fs.Int("v", 0, "version to force")
		_ = fs.Parse(flag.Args()[1:])
		if err := m.Force(*v); err != nil {
			return err
		}
		fmt.Printf("forced version=%d\n", *v)
	}
	return nil
}

func createCmd(dir string) error {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	name := fs.String("name", "", "migration name (snake_case)")
	_ = fs.Parse(flag.Args()[1:])
	if *name == "" {
		return fmt.Errorf("missing -name")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	ts := time.Now().UTC().Format("20060102150405")
	up := filepath.Join(dir, fmt.Sprintf("%s_%s.up.sql", ts, *name))
	down := filepath.Join(dir, fmt.Sprintf("%s_%s.down.sql", ts, *name))
	if err := os.WriteFile(up, []byte("-- "+*name+" up\n"), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(down, []byte("-- "+*name+" down\n"), 0o644); err != nil {
		return err
	}
	fmt.Println(up)
	fmt.Println(down)
	return nil
}

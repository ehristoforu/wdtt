package panel

import "strings"

func migratePanelDBV15() error {
	_, err := panelDB.Exec(`ALTER TABLE wdtt_inbound ADD COLUMN raw_enable INTEGER NOT NULL DEFAULT 1`)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "duplicate column") {
		return err
	}
	_, err = panelDB.Exec(`ALTER TABLE wdtt_inbound ADD COLUMN raw_direct_port INTEGER NOT NULL DEFAULT 0`)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "duplicate column") {
		return err
	}
	return nil
}

// migratePanelDBV16 raises the inbound user capacity for existing installs.
func migratePanelDBV16() error {
	_, err := panelDB.Exec(`UPDATE wdtt_inbound SET max_users = 249`)
	return err
}

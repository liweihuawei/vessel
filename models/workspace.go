package models

import (
	// "time"
	//	"github.com/huawei-openlab/newdb/orm"
	//	"github.com/ngaut/log"
)

type Workspace struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Actived     bool      `json:"actived"`
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	Memo        string    `json:"memo"`
}

func (ws *Workspace) Create(name, description string) (int64, error) {
	/*	o := orm.NewOrm()
		w := Workspace{Name: name, Description: description, Actived: true}

		if err := o.Begin(); err != nil {
			log.Errorf("Transcation error: %s", err.Error())

			return 0, err
		} else {
			if id, e := o.Insert(&w); e != nil {
				log.Errorf("Create workspace error: %s", e.Error())

				o.Rollback()
				return 0, e
			} else {
				log.Infof("Create workspace successfully, id is: %d", id)

				o.Commit()
				return id, nil
			}
		}
	*/
	return 0, nil
}

func (ws *Workspace) Put(id int64, name, description string) error {
	/*
		o := orm.NewOrm()
		w := Workspace{Id: id, Actived: true}

		if err := o.Read(&w, "Id", "Actived"); err != nil {
			log.Errorf("Get workspace %d error: %s", id, err.Error())

			return err
		} else {
			if err := o.Begin(); err != nil {
				log.Errorf("Transcation error: %s", err.Error())

				o.Rollback()
				return err
			} else {
				w.Name = name
				w.Description = description

				if _, err := o.Update(&w, "Name", "Description"); err != nil {
					log.Errorf("Put workspace %d error: %s", id, err.Error())

					o.Rollback()
					return err
				} else {
					log.Infof("Put workspace successfully: %d", id)

					o.Commit()
					return nil
				}
			}
		}
	*/
	return nil
}

func (ws *Workspace) Get(id int64) (*Workspace, error) {
	/*
		o := orm.NewOrm()
		w := Workspace{Id: id, Actived: true}

		if err := o.Read(&w, "Id", "Actived"); err != nil {
			log.Errorf("Get workspace %d error: %s", id, err.Error())

			return w, err
		} else {
			return w, nil
		}
	*/
	return nil, nil
}

func (ws *Workspace) Delete(id int64) error {
	/*
		o := orm.NewOrm()
		w := Workspace{Id: id}

		if err := o.Read(&w, "Id"); err != nil {
			log.Errorf("Get workspace %d error: %s", id, err.Error())

			return err
		} else {
			if err := o.Begin(); err != nil {
				log.Errorf("Transcation error: %s", err.Error())

				o.Rollback()
				return err
			} else {
				if _, err := o.Delete(&w); err != nil {
					log.Errorf("Delete workspace %d error: %s", id, err.Error())

					o.Rollback()
					return err
				} else {
					log.Infof("Delete workspace %d successfully", id)

					pjs := []*Project{}

					if _, err := o.QueryTable("project").Filter("workspace_id", id).All(&pjs); err != nil {
						log.Errorf("Get all projects of workspace %d error: %s", id, err.Error())

						o.Rollback()
						return err
					}

					for _, p := range pjs {
						project := Project{}

						if err := project.Delete(p.Id); err != nil {
							o.Rollback()
							return err
						}
					}

					o.Commit()
					return nil
				}
			}
		}
	*/
	return nil
}

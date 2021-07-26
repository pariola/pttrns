package proxy

import (
	"errors"
)

type User struct {
	Id int32
}

type UserFinder interface {
	Find(id int32) (*User, error)
}

type UserList []*User

func (l UserList) Find(id int32) (*User, error) {

	for _, user := range l {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("not found")
}

type UserListProxy struct {
	capacity            int
	cache               UserList
	database            UserList
	lastSearchFromCache bool
}

func (p *UserListProxy) addToStack(u *User) {
	if len(p.cache) >= p.capacity {
		p.cache = append(p.cache[1:], u)
	} else {
		p.cache = append(p.cache, u)
	}
}

func (p *UserListProxy) Find(id int32) (*User, error) {

	user, err := p.cache.Find(id)

	// in-cache
	if err == nil {
		p.lastSearchFromCache = true
		return user, nil
	}

	user, err = p.database.Find(id)

	if err != nil {
		return nil, err
	}

	p.lastSearchFromCache = false
	p.addToStack(user)

	return user, nil
}

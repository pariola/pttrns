package proxy

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {

	database := UserList{}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10000; i++ {
		database = append(database, &User{Id: rand.Int31()})
	}

	testIds := [3]int32{database[100].Id, database[400].Id, database[900].Id}

	proxy := UserListProxy{
		capacity: 2,
		cache:    UserList{},
		database: database,
	}

	// test search
	user, err := proxy.Find(testIds[0])
	require.Nil(t, err)
	assert.Equal(t, user.Id, testIds[0], "search id & returned user id don't match")
	assert.Equal(t, proxy.lastSearchFromCache, false, "search should not use cache")

	// test cache
	user, err = proxy.Find(testIds[0])
	assert.Equal(t, proxy.lastSearchFromCache, true, "search should use cache")

	// test cache capacity
	user, err = proxy.Find(testIds[1])
	assert.Equal(t, proxy.lastSearchFromCache, false, "search should not use cache")

	user, err = proxy.Find(testIds[2])
	assert.Equal(t, proxy.lastSearchFromCache, false, "search should not use cache")

	// test cache - cache shouldn't be used since capacity is set to 2
	user, err = proxy.Find(testIds[0])
	assert.Equal(t, proxy.lastSearchFromCache, false, "search should not use cache")
}

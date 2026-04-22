# interfacebloat

<instructions>
Interfacebloat detects interfaces with too many methods (default threshold: 10). Large interfaces violate the Interface Segregation Principle — callers depend on methods they don't use, making implementations harder to maintain and test.

Split large interfaces into smaller, focused ones composed together. Consumers should depend only on the methods they need.
</instructions>

<examples>
## Bad
```go
type Repository interface {
    Find(id int) (*User, error)
    FindAll() ([]User, error)
    Save(u *User) error
    Delete(id int) error
    FindPosts(userID int) ([]Post, error)
    SavePost(p *Post) error
    DeletePost(id int) error
    FindComments(postID int) ([]Comment, error)
    SaveComment(c *Comment) error
    DeleteComment(id int) error
}
```

## Good
```go
type UserReader interface {
    Find(id int) (*User, error)
    FindAll() ([]User, error)
}

type UserWriter interface {
    Save(u *User) error
    Delete(id int) error
}

type PostRepository interface {
    FindPosts(userID int) ([]Post, error)
    SavePost(p *Post) error
    DeletePost(id int) error
}
```
</examples>

<patterns>
- Split mega-interfaces into entity-specific interfaces for each CRUD concern
- Separate service interfaces into query, command, and event interfaces
- Decompose client interfaces into focused sub-interfaces by API concern
</patterns>

<related>
funlen, gocyclo, revive
</related>

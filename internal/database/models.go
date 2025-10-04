package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull,unique"`
	EmailVerified bool      `bun:"email_verified,default:false"`
	Image         *string   `bun:"image"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type Account struct {
	bun.BaseModel `bun:"table:account,alias:a"`

	ID                    uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID                uuid.UUID  `bun:"user_id,notnull"`
	AccountID             string     `bun:"account_id,notnull"`
	ProviderID            string     `bun:"provider_id,notnull"`
	AccessToken           *string    `bun:"access_token"`
	RefreshToken          *string    `bun:"refresh_token"`
	AccessTokenExpiresAt  *time.Time `bun:"access_token_expires_at"`
	RefreshTokenExpiresAt *time.Time `bun:"refresh_token_expires_at"`
	Scope                 *string    `bun:"scope"`
	IDToken               *string    `bun:"id_token"`
	CreatedAt             time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt             time.Time  `bun:"updated_at,notnull,default:current_timestamp"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

type Session struct {
	bun.BaseModel `bun:"table:session,alias:s"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID    uuid.UUID `bun:"user_id,notnull"`
	Token     string    `bun:"token,notnull,unique"`
	ExpiresAt time.Time `bun:"expires_at,notnull"`
	IPAddress *string   `bun:"ip_address"`
	UserAgent *string   `bun:"user_agent"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

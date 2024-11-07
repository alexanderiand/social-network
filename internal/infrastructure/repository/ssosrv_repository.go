package repository

// SSOServiceRepository
type SSOServiceRepository interface {
	UserRepository
	// TODO: AuthRepository
	// TODO: ProfileRepository
}

// UserRepository
type UserRepository interface {
	// TODO: UserCreator
	// TODO: UserProvider
	// TODO: UserUpdater
	// TODO: UserDeleter

	// TODO: UserNotifier // maybe in transport layer
}

// AuthRepository

// ProfileRepository

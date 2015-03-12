package gobs

type userData struct {
    data
}

type usersRegister struct {
    Username string
    Password string
    DataItem interface{}
}

type socialIdentity struct {
    Provider string
    Token string
    TokenSecret string
}

func (u userData) LoginWithFacebook(accessToken string) error {
    fbLogin := socialIdentity{Provider: "Facebook", Token: accessToken}
    return u.loginWithProvider(fbLogin)
}

func (u userData) LinkWithFacebook(accessToken string, userId string) error {
    fbLink := socialIdentity{Provider: "Facebook", Token: accessToken}

    return u.linkWithProvider(fbLink, userId)
}

func (u userData) UnlinkFromFacebook(userId string) error {
    return u.unlinkFromProvider(userId, "Facebook")
}

func (u userData) LoginWithGoogle(accessToken string) error {
    googleLogin := socialIdentity{Provider: "Google", Token: accessToken}

    return u.loginWithProvider(googleLogin)
}

func (u userData) LinkWithGoogle(accessToken string, userId string) error {
    googleLink := socialIdentity{Provider: "Google", Token: accessToken}
    return u.linkWithProvider(googleLink, userId)
}

func (u userData) UnlinkFromGoogle(userId string) error {
    return u.unlinkFromProvider(userId, "Google")
}

func (u userData) LoginWithLiveID(accessToken string) error {
    liveIdLogin := socialIdentity{Provider: "LiveID", Token: accessToken}

    return u.loginWithProvider(liveIdLogin)
}

func (u userData) LinkWithLiveID(accessToken string, userId string) error {
    liveIdLink := socialIdentity{Provider: "LiveID", Token: accessToken}
    return u.linkWithProvider(liveIdLink, userId)
}

func (u userData) UnlinkFromLiveID(userId string) error {
    return u.unlinkFromProvider(userId, "LiveID")
}

func (u userData) LoginWithADFS(accessToken string) error {
    adfsLogin := socialIdentity{Provider: "ADFS", Token: accessToken}

    return u.loginWithProvider(adfsLogin)
}

func (u userData) LinkWithADFS(accessToken string, userId string) error {
    adfsLink := socialIdentity{Provider: "ADFS", Token: accessToken}
    return u.linkWithProvider(adfsLink, userId)
}


func (u userData) UnlinkFromADFS(userId string) error {
    return u.unlinkFromProvider(userId, "ADFS")
}
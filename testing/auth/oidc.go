package auth

// OAuthTestClient creates a new http client handling OAuth automatically.
// Returned is the new HTTP Client, OIDC URI and a close function.
func OAuthTestClient(subject string, audience string) (*http.Client, string, func()) {
	issuer, closer := testHelperOIDCProvider(TestPrivRSAKey1ID, TestPrivRSAKey2ID)

	ctx := context.Background()

	var audiences jwt.Audience

	if audience != "" {
		audiences = append(audiences, audience)
	}

	authClaim := jwt.Claims{
		Issuer:    issuer,
		Subject:   subject,
		Audience:  audiences,
		NotBefore: jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
	}

	signer := testHelperMustMakeSigner(jose.RS256, TestPrivRSAKey1ID, TestPrivRSAKey1)
	rawToken := testHelperGetToken(signer, authClaim, "scope", "test")

	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: rawToken,
	})), issuer, closer
}

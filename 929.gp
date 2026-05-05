func numUniqueEmails(emails []string) int {
    var result int
    set := make(map[string]struct{}, len(emails))
    for _, email := range emails {
        local, domain, ok := strings.Cut(email, "@")
        if !ok || !strings.HasSuffix(domain, ".com") || len(domain) == 4 {
            continue
        }

        local, _, _ = strings.Cut(strings.ReplaceAll(local, ".", ""), "+")
        email = local + "@" + domain
        if _, ok = set[email]; !ok {
            result++
            set[email] = struct{}{}
        }
    }
    return result
}

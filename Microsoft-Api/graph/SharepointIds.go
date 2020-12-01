package graph

import(
  "net/url"
)

type SharePointIds struct{
    ListId           string  `json "listId"`
    ListItemId       string  `json "listItemId"`
    ListItemUniqueId string  `json "listItemUniqueId"`
    SiteId           string  `json "siteId"`
    SiteUrl          url.URL `json "siteUrl"`
    TenantId         string  `json "tenantId"`
    WebId            string  `json "webId"`
}

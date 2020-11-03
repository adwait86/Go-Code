package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
)

func main() {
	runGitLabParsingTest()

	runGitHubParsingTest()

}

func runGitLabParsingTest() {
	body := strings.NewReader(getGitLabPayloadForPull())
	r, err := http.NewRequest("POST", "/", body)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Gitlab-Event", "Merge Request Hook")
	r.Header.Set("X-Gitlab-Token", "secret")
	r.Header.Set("User-Agent", "GitLab/13.6.0-pre")

	if err != nil {
		fmt.Println("fail")
	}

	c := gitlab.NewDefault()
	h, err := c.Webhooks.Parse(r, func(scm.Webhook) (string, error) {
		return "secret", nil
	})

	fmt.Println(h)
}

func runGitHubParsingTest() {
	body := strings.NewReader(getGithubPayloadForPull())
	r, err := http.NewRequest("POST", "/", body)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-GitHub-Event", "pull_request")

	if err != nil {
		fmt.Println("fail")
	}

	c := github.NewDefault()
	h, err := c.Webhooks.Parse(r, func(scm.Webhook) (string, error) {
		return "", nil
	})

	fmt.Println(h)
}

func getGithubPayloadForPull() string {
	return `{
		"action": "synchronize",
		"number": 1,
		"pull_request": {
		  "url": "https://api.github.com/repos/adwait86/test2/pulls/1",
		  "id": 509353795,
		  "node_id": "MDExOlB1bGxSZXF1ZXN0NTA5MzUzNzk1",
		  "html_url": "https://github.com/adwait86/test2/pull/1",
		  "diff_url": "https://github.com/adwait86/test2/pull/1.diff",
		  "patch_url": "https://github.com/adwait86/test2/pull/1.patch",
		  "issue_url": "https://api.github.com/repos/adwait86/test2/issues/1",
		  "number": 1,
		  "state": "open",
		  "locked": false,
		  "title": "c",
		  "user": {
			"login": "adwait86",
			"id": 13191390,
			"node_id": "MDQ6VXNlcjEzMTkxMzkw",
			"avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/adwait86",
			"html_url": "https://github.com/adwait86",
			"followers_url": "https://api.github.com/users/adwait86/followers",
			"following_url": "https://api.github.com/users/adwait86/following{/other_user}",
			"gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
			"organizations_url": "https://api.github.com/users/adwait86/orgs",
			"repos_url": "https://api.github.com/users/adwait86/repos",
			"events_url": "https://api.github.com/users/adwait86/events{/privacy}",
			"received_events_url": "https://api.github.com/users/adwait86/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "created_at": "2020-10-24T06:39:56Z",
		  "updated_at": "2020-11-03T00:04:11Z",
		  "closed_at": null,
		  "merged_at": null,
		  "merge_commit_sha": "720bc8c91816a28a495b93a28ee5cb1cd3d8f5bf",
		  "assignee": null,
		  "assignees": [
	  
		  ],
		  "requested_reviewers": [
	  
		  ],
		  "requested_teams": [
	  
		  ],
		  "labels": [
	  
		  ],
		  "milestone": null,
		  "draft": false,
		  "commits_url": "https://api.github.com/repos/adwait86/test2/pulls/1/commits",
		  "review_comments_url": "https://api.github.com/repos/adwait86/test2/pulls/1/comments",
		  "review_comment_url": "https://api.github.com/repos/adwait86/test2/pulls/comments{/number}",
		  "comments_url": "https://api.github.com/repos/adwait86/test2/issues/1/comments",
		  "statuses_url": "https://api.github.com/repos/adwait86/test2/statuses/da0b8fca101b1407bdfc9a14aa2f78c6aa06f71e",
		  "head": {
			"label": "adwait86:test",
			"ref": "test",
			"sha": "da0b8fca101b1407bdfc9a14aa2f78c6aa06f71e",
			"user": {
			  "login": "adwait86",
			  "id": 13191390,
			  "node_id": "MDQ6VXNlcjEzMTkxMzkw",
			  "avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/adwait86",
			  "html_url": "https://github.com/adwait86",
			  "followers_url": "https://api.github.com/users/adwait86/followers",
			  "following_url": "https://api.github.com/users/adwait86/following{/other_user}",
			  "gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
			  "organizations_url": "https://api.github.com/users/adwait86/orgs",
			  "repos_url": "https://api.github.com/users/adwait86/repos",
			  "events_url": "https://api.github.com/users/adwait86/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/adwait86/received_events",
			  "type": "User",
			  "site_admin": false
			},
			"repo": {
			  "id": 42551937,
			  "node_id": "MDEwOlJlcG9zaXRvcnk0MjU1MTkzNw==",
			  "name": "test2",
			  "full_name": "adwait86/test2",
			  "private": false,
			  "owner": {
				"login": "adwait86",
				"id": 13191390,
				"node_id": "MDQ6VXNlcjEzMTkxMzkw",
				"avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
				"gravatar_id": "",
				"url": "https://api.github.com/users/adwait86",
				"html_url": "https://github.com/adwait86",
				"followers_url": "https://api.github.com/users/adwait86/followers",
				"following_url": "https://api.github.com/users/adwait86/following{/other_user}",
				"gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
				"organizations_url": "https://api.github.com/users/adwait86/orgs",
				"repos_url": "https://api.github.com/users/adwait86/repos",
				"events_url": "https://api.github.com/users/adwait86/events{/privacy}",
				"received_events_url": "https://api.github.com/users/adwait86/received_events",
				"type": "User",
				"site_admin": false
			  },
			  "html_url": "https://github.com/adwait86/test2",
			  "description": null,
			  "fork": false,
			  "url": "https://api.github.com/repos/adwait86/test2",
			  "forks_url": "https://api.github.com/repos/adwait86/test2/forks",
			  "keys_url": "https://api.github.com/repos/adwait86/test2/keys{/key_id}",
			  "collaborators_url": "https://api.github.com/repos/adwait86/test2/collaborators{/collaborator}",
			  "teams_url": "https://api.github.com/repos/adwait86/test2/teams",
			  "hooks_url": "https://api.github.com/repos/adwait86/test2/hooks",
			  "issue_events_url": "https://api.github.com/repos/adwait86/test2/issues/events{/number}",
			  "events_url": "https://api.github.com/repos/adwait86/test2/events",
			  "assignees_url": "https://api.github.com/repos/adwait86/test2/assignees{/user}",
			  "branches_url": "https://api.github.com/repos/adwait86/test2/branches{/branch}",
			  "tags_url": "https://api.github.com/repos/adwait86/test2/tags",
			  "blobs_url": "https://api.github.com/repos/adwait86/test2/git/blobs{/sha}",
			  "git_tags_url": "https://api.github.com/repos/adwait86/test2/git/tags{/sha}",
			  "git_refs_url": "https://api.github.com/repos/adwait86/test2/git/refs{/sha}",
			  "trees_url": "https://api.github.com/repos/adwait86/test2/git/trees{/sha}",
			  "statuses_url": "https://api.github.com/repos/adwait86/test2/statuses/{sha}",
			  "languages_url": "https://api.github.com/repos/adwait86/test2/languages",
			  "stargazers_url": "https://api.github.com/repos/adwait86/test2/stargazers",
			  "contributors_url": "https://api.github.com/repos/adwait86/test2/contributors",
			  "subscribers_url": "https://api.github.com/repos/adwait86/test2/subscribers",
			  "subscription_url": "https://api.github.com/repos/adwait86/test2/subscription",
			  "commits_url": "https://api.github.com/repos/adwait86/test2/commits{/sha}",
			  "git_commits_url": "https://api.github.com/repos/adwait86/test2/git/commits{/sha}",
			  "comments_url": "https://api.github.com/repos/adwait86/test2/comments{/number}",
			  "issue_comment_url": "https://api.github.com/repos/adwait86/test2/issues/comments{/number}",
			  "contents_url": "https://api.github.com/repos/adwait86/test2/contents/{+path}",
			  "compare_url": "https://api.github.com/repos/adwait86/test2/compare/{base}...{head}",
			  "merges_url": "https://api.github.com/repos/adwait86/test2/merges",
			  "archive_url": "https://api.github.com/repos/adwait86/test2/{archive_format}{/ref}",
			  "downloads_url": "https://api.github.com/repos/adwait86/test2/downloads",
			  "issues_url": "https://api.github.com/repos/adwait86/test2/issues{/number}",
			  "pulls_url": "https://api.github.com/repos/adwait86/test2/pulls{/number}",
			  "milestones_url": "https://api.github.com/repos/adwait86/test2/milestones{/number}",
			  "notifications_url": "https://api.github.com/repos/adwait86/test2/notifications{?since,all,participating}",
			  "labels_url": "https://api.github.com/repos/adwait86/test2/labels{/name}",
			  "releases_url": "https://api.github.com/repos/adwait86/test2/releases{/id}",
			  "deployments_url": "https://api.github.com/repos/adwait86/test2/deployments",
			  "created_at": "2015-09-15T23:04:38Z",
			  "updated_at": "2020-10-24T06:34:58Z",
			  "pushed_at": "2020-11-03T00:04:10Z",
			  "git_url": "git://github.com/adwait86/test2.git",
			  "ssh_url": "git@github.com:adwait86/test2.git",
			  "clone_url": "https://github.com/adwait86/test2.git",
			  "svn_url": "https://github.com/adwait86/test2",
			  "homepage": null,
			  "size": 2,
			  "stargazers_count": 0,
			  "watchers_count": 0,
			  "language": null,
			  "has_issues": true,
			  "has_projects": true,
			  "has_downloads": true,
			  "has_wiki": true,
			  "has_pages": false,
			  "forks_count": 0,
			  "mirror_url": null,
			  "archived": false,
			  "disabled": false,
			  "open_issues_count": 1,
			  "license": null,
			  "forks": 0,
			  "open_issues": 1,
			  "watchers": 0,
			  "default_branch": "master",
			  "allow_squash_merge": true,
			  "allow_merge_commit": true,
			  "allow_rebase_merge": true,
			  "delete_branch_on_merge": false
			}
		  },
		  "base": {
			"label": "adwait86:master",
			"ref": "master",
			"sha": "4af2388e5dc7d57e28b5326173a3a6bbc8bfe111",
			"user": {
			  "login": "adwait86",
			  "id": 13191390,
			  "node_id": "MDQ6VXNlcjEzMTkxMzkw",
			  "avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/adwait86",
			  "html_url": "https://github.com/adwait86",
			  "followers_url": "https://api.github.com/users/adwait86/followers",
			  "following_url": "https://api.github.com/users/adwait86/following{/other_user}",
			  "gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
			  "organizations_url": "https://api.github.com/users/adwait86/orgs",
			  "repos_url": "https://api.github.com/users/adwait86/repos",
			  "events_url": "https://api.github.com/users/adwait86/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/adwait86/received_events",
			  "type": "User",
			  "site_admin": false
			},
			"repo": {
			  "id": 42551937,
			  "node_id": "MDEwOlJlcG9zaXRvcnk0MjU1MTkzNw==",
			  "name": "test2",
			  "full_name": "adwait86/test2",
			  "private": false,
			  "owner": {
				"login": "adwait86",
				"id": 13191390,
				"node_id": "MDQ6VXNlcjEzMTkxMzkw",
				"avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
				"gravatar_id": "",
				"url": "https://api.github.com/users/adwait86",
				"html_url": "https://github.com/adwait86",
				"followers_url": "https://api.github.com/users/adwait86/followers",
				"following_url": "https://api.github.com/users/adwait86/following{/other_user}",
				"gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
				"organizations_url": "https://api.github.com/users/adwait86/orgs",
				"repos_url": "https://api.github.com/users/adwait86/repos",
				"events_url": "https://api.github.com/users/adwait86/events{/privacy}",
				"received_events_url": "https://api.github.com/users/adwait86/received_events",
				"type": "User",
				"site_admin": false
			  },
			  "html_url": "https://github.com/adwait86/test2",
			  "description": null,
			  "fork": false,
			  "url": "https://api.github.com/repos/adwait86/test2",
			  "forks_url": "https://api.github.com/repos/adwait86/test2/forks",
			  "keys_url": "https://api.github.com/repos/adwait86/test2/keys{/key_id}",
			  "collaborators_url": "https://api.github.com/repos/adwait86/test2/collaborators{/collaborator}",
			  "teams_url": "https://api.github.com/repos/adwait86/test2/teams",
			  "hooks_url": "https://api.github.com/repos/adwait86/test2/hooks",
			  "issue_events_url": "https://api.github.com/repos/adwait86/test2/issues/events{/number}",
			  "events_url": "https://api.github.com/repos/adwait86/test2/events",
			  "assignees_url": "https://api.github.com/repos/adwait86/test2/assignees{/user}",
			  "branches_url": "https://api.github.com/repos/adwait86/test2/branches{/branch}",
			  "tags_url": "https://api.github.com/repos/adwait86/test2/tags",
			  "blobs_url": "https://api.github.com/repos/adwait86/test2/git/blobs{/sha}",
			  "git_tags_url": "https://api.github.com/repos/adwait86/test2/git/tags{/sha}",
			  "git_refs_url": "https://api.github.com/repos/adwait86/test2/git/refs{/sha}",
			  "trees_url": "https://api.github.com/repos/adwait86/test2/git/trees{/sha}",
			  "statuses_url": "https://api.github.com/repos/adwait86/test2/statuses/{sha}",
			  "languages_url": "https://api.github.com/repos/adwait86/test2/languages",
			  "stargazers_url": "https://api.github.com/repos/adwait86/test2/stargazers",
			  "contributors_url": "https://api.github.com/repos/adwait86/test2/contributors",
			  "subscribers_url": "https://api.github.com/repos/adwait86/test2/subscribers",
			  "subscription_url": "https://api.github.com/repos/adwait86/test2/subscription",
			  "commits_url": "https://api.github.com/repos/adwait86/test2/commits{/sha}",
			  "git_commits_url": "https://api.github.com/repos/adwait86/test2/git/commits{/sha}",
			  "comments_url": "https://api.github.com/repos/adwait86/test2/comments{/number}",
			  "issue_comment_url": "https://api.github.com/repos/adwait86/test2/issues/comments{/number}",
			  "contents_url": "https://api.github.com/repos/adwait86/test2/contents/{+path}",
			  "compare_url": "https://api.github.com/repos/adwait86/test2/compare/{base}...{head}",
			  "merges_url": "https://api.github.com/repos/adwait86/test2/merges",
			  "archive_url": "https://api.github.com/repos/adwait86/test2/{archive_format}{/ref}",
			  "downloads_url": "https://api.github.com/repos/adwait86/test2/downloads",
			  "issues_url": "https://api.github.com/repos/adwait86/test2/issues{/number}",
			  "pulls_url": "https://api.github.com/repos/adwait86/test2/pulls{/number}",
			  "milestones_url": "https://api.github.com/repos/adwait86/test2/milestones{/number}",
			  "notifications_url": "https://api.github.com/repos/adwait86/test2/notifications{?since,all,participating}",
			  "labels_url": "https://api.github.com/repos/adwait86/test2/labels{/name}",
			  "releases_url": "https://api.github.com/repos/adwait86/test2/releases{/id}",
			  "deployments_url": "https://api.github.com/repos/adwait86/test2/deployments",
			  "created_at": "2015-09-15T23:04:38Z",
			  "updated_at": "2020-10-24T06:34:58Z",
			  "pushed_at": "2020-11-03T00:04:10Z",
			  "git_url": "git://github.com/adwait86/test2.git",
			  "ssh_url": "git@github.com:adwait86/test2.git",
			  "clone_url": "https://github.com/adwait86/test2.git",
			  "svn_url": "https://github.com/adwait86/test2",
			  "homepage": null,
			  "size": 2,
			  "stargazers_count": 0,
			  "watchers_count": 0,
			  "language": null,
			  "has_issues": true,
			  "has_projects": true,
			  "has_downloads": true,
			  "has_wiki": true,
			  "has_pages": false,
			  "forks_count": 0,
			  "mirror_url": null,
			  "archived": false,
			  "disabled": false,
			  "open_issues_count": 1,
			  "license": null,
			  "forks": 0,
			  "open_issues": 1,
			  "watchers": 0,
			  "default_branch": "master",
			  "allow_squash_merge": true,
			  "allow_merge_commit": true,
			  "allow_rebase_merge": true,
			  "delete_branch_on_merge": false
			}
		  },
		  "_links": {
			"self": {
			  "href": "https://api.github.com/repos/adwait86/test2/pulls/1"
			},
			"html": {
			  "href": "https://github.com/adwait86/test2/pull/1"
			},
			"issue": {
			  "href": "https://api.github.com/repos/adwait86/test2/issues/1"
			},
			"comments": {
			  "href": "https://api.github.com/repos/adwait86/test2/issues/1/comments"
			},
			"review_comments": {
			  "href": "https://api.github.com/repos/adwait86/test2/pulls/1/comments"
			},
			"review_comment": {
			  "href": "https://api.github.com/repos/adwait86/test2/pulls/comments{/number}"
			},
			"commits": {
			  "href": "https://api.github.com/repos/adwait86/test2/pulls/1/commits"
			},
			"statuses": {
			  "href": "https://api.github.com/repos/adwait86/test2/statuses/da0b8fca101b1407bdfc9a14aa2f78c6aa06f71e"
			}
		  },
		  "author_association": "OWNER",
		  "active_lock_reason": null,
		  "merged": false,
		  "mergeable": null,
		  "rebaseable": null,
		  "mergeable_state": "unknown",
		  "merged_by": null,
		  "comments": 0,
		  "review_comments": 0,
		  "maintainer_can_modify": false,
		  "commits": 3,
		  "additions": 3,
		  "deletions": 0,
		  "changed_files": 3
		},
		"before": "7f8519e9688cfdcad00d508bdeef35ba5ef3d09f",
		"after": "da0b8fca101b1407bdfc9a14aa2f78c6aa06f71e",
		"repository": {
		  "id": 42551937,
		  "node_id": "MDEwOlJlcG9zaXRvcnk0MjU1MTkzNw==",
		  "name": "test2",
		  "full_name": "adwait86/test2",
		  "private": false,
		  "owner": {
			"login": "adwait86",
			"id": 13191390,
			"node_id": "MDQ6VXNlcjEzMTkxMzkw",
			"avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/adwait86",
			"html_url": "https://github.com/adwait86",
			"followers_url": "https://api.github.com/users/adwait86/followers",
			"following_url": "https://api.github.com/users/adwait86/following{/other_user}",
			"gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
			"organizations_url": "https://api.github.com/users/adwait86/orgs",
			"repos_url": "https://api.github.com/users/adwait86/repos",
			"events_url": "https://api.github.com/users/adwait86/events{/privacy}",
			"received_events_url": "https://api.github.com/users/adwait86/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "html_url": "https://github.com/adwait86/test2",
		  "description": null,
		  "fork": false,
		  "url": "https://api.github.com/repos/adwait86/test2",
		  "forks_url": "https://api.github.com/repos/adwait86/test2/forks",
		  "keys_url": "https://api.github.com/repos/adwait86/test2/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/adwait86/test2/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/adwait86/test2/teams",
		  "hooks_url": "https://api.github.com/repos/adwait86/test2/hooks",
		  "issue_events_url": "https://api.github.com/repos/adwait86/test2/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/adwait86/test2/events",
		  "assignees_url": "https://api.github.com/repos/adwait86/test2/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/adwait86/test2/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/adwait86/test2/tags",
		  "blobs_url": "https://api.github.com/repos/adwait86/test2/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/adwait86/test2/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/adwait86/test2/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/adwait86/test2/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/adwait86/test2/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/adwait86/test2/languages",
		  "stargazers_url": "https://api.github.com/repos/adwait86/test2/stargazers",
		  "contributors_url": "https://api.github.com/repos/adwait86/test2/contributors",
		  "subscribers_url": "https://api.github.com/repos/adwait86/test2/subscribers",
		  "subscription_url": "https://api.github.com/repos/adwait86/test2/subscription",
		  "commits_url": "https://api.github.com/repos/adwait86/test2/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/adwait86/test2/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/adwait86/test2/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/adwait86/test2/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/adwait86/test2/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/adwait86/test2/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/adwait86/test2/merges",
		  "archive_url": "https://api.github.com/repos/adwait86/test2/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/adwait86/test2/downloads",
		  "issues_url": "https://api.github.com/repos/adwait86/test2/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/adwait86/test2/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/adwait86/test2/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/adwait86/test2/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/adwait86/test2/labels{/name}",
		  "releases_url": "https://api.github.com/repos/adwait86/test2/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/adwait86/test2/deployments",
		  "created_at": "2015-09-15T23:04:38Z",
		  "updated_at": "2020-10-24T06:34:58Z",
		  "pushed_at": "2020-11-03T00:04:10Z",
		  "git_url": "git://github.com/adwait86/test2.git",
		  "ssh_url": "git@github.com:adwait86/test2.git",
		  "clone_url": "https://github.com/adwait86/test2.git",
		  "svn_url": "https://github.com/adwait86/test2",
		  "homepage": null,
		  "size": 2,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": true,
		  "has_pages": false,
		  "forks_count": 0,
		  "mirror_url": null,
		  "archived": false,
		  "disabled": false,
		  "open_issues_count": 1,
		  "license": null,
		  "forks": 0,
		  "open_issues": 1,
		  "watchers": 0,
		  "default_branch": "master"
		},
		"sender": {
		  "login": "adwait86",
		  "id": 13191390,
		  "node_id": "MDQ6VXNlcjEzMTkxMzkw",
		  "avatar_url": "https://avatars0.githubusercontent.com/u/13191390?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/adwait86",
		  "html_url": "https://github.com/adwait86",
		  "followers_url": "https://api.github.com/users/adwait86/followers",
		  "following_url": "https://api.github.com/users/adwait86/following{/other_user}",
		  "gists_url": "https://api.github.com/users/adwait86/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/adwait86/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/adwait86/subscriptions",
		  "organizations_url": "https://api.github.com/users/adwait86/orgs",
		  "repos_url": "https://api.github.com/users/adwait86/repos",
		  "events_url": "https://api.github.com/users/adwait86/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/adwait86/received_events",
		  "type": "User",
		  "site_admin": false
		}
	  }`
}
func getGitLabPayloadForPull() string {
	return `{
		"object_kind": "merge_request",
		"event_type": "merge_request",
		"user": {
		  "name": "Adwait Bhandare",
		  "username": "bhandare.adwait",
		  "avatar_url": "https://secure.gravatar.com/avatar/8e4c3976618e1b6e74d01d12c4ace96c?s=80&d=identicon",
		  "email": "bhandare.adwait@gmail.com"
		},
		"project": {
		  "id": 21699507,
		  "name": "test",
		  "description": "",
		  "web_url": "https://gitlab.com/adwaittest/test",
		  "avatar_url": null,
		  "git_ssh_url": "git@gitlab.com:adwaittest/test.git",
		  "git_http_url": "https://gitlab.com/adwaittest/test.git",
		  "namespace": "adwaitTest",
		  "visibility_level": 0,
		  "path_with_namespace": "adwaittest/test",
		  "default_branch": "master",
		  "ci_config_path": "",
		  "homepage": "https://gitlab.com/adwaittest/test",
		  "url": "git@gitlab.com:adwaittest/test.git",
		  "ssh_url": "git@gitlab.com:adwaittest/test.git",
		  "http_url": "https://gitlab.com/adwaittest/test.git"
		},
		"object_attributes": {
		  "assignee_id": null,
		  "author_id": 7341710,
		  "created_at": "2020-11-02 23:35:31 UTC",
		  "description": "",
		  "head_pipeline_id": null,
		  "id": 76594814,
		  "iid": 1,
		  "last_edited_at": null,
		  "last_edited_by_id": null,
		  "merge_commit_sha": null,
		  "merge_error": null,
		  "merge_params": {
			"force_remove_source_branch": "1"
		  },
		  "merge_status": "unchecked",
		  "merge_user_id": null,
		  "merge_when_pipeline_succeeds": false,
		  "milestone_id": null,
		  "source_branch": "stage",
		  "source_project_id": 21699507,
		  "state_id": 1,
		  "target_branch": "master",
		  "target_project_id": 21699507,
		  "time_estimate": 0,
		  "title": "Add new file 5",
		  "updated_at": "2020-11-02 23:35:31 UTC",
		  "updated_by_id": null,
		  "url": "https://gitlab.com/adwaittest/test/-/merge_requests/1",
		  "source": {
			"id": 21699507,
			"name": "test",
			"description": "",
			"web_url": "https://gitlab.com/adwaittest/test",
			"avatar_url": null,
			"git_ssh_url": "git@gitlab.com:adwaittest/test.git",
			"git_http_url": "https://gitlab.com/adwaittest/test.git",
			"namespace": "adwaitTest",
			"visibility_level": 0,
			"path_with_namespace": "adwaittest/test",
			"default_branch": "master",
			"ci_config_path": "",
			"homepage": "https://gitlab.com/adwaittest/test",
			"url": "git@gitlab.com:adwaittest/test.git",
			"ssh_url": "git@gitlab.com:adwaittest/test.git",
			"http_url": "https://gitlab.com/adwaittest/test.git"
		  },
		  "target": {
			"id": 21699507,
			"name": "test",
			"description": "",
			"web_url": "https://gitlab.com/adwaittest/test",
			"avatar_url": null,
			"git_ssh_url": "git@gitlab.com:adwaittest/test.git",
			"git_http_url": "https://gitlab.com/adwaittest/test.git",
			"namespace": "adwaitTest",
			"visibility_level": 0,
			"path_with_namespace": "adwaittest/test",
			"default_branch": "master",
			"ci_config_path": "",
			"homepage": "https://gitlab.com/adwaittest/test",
			"url": "git@gitlab.com:adwaittest/test.git",
			"ssh_url": "git@gitlab.com:adwaittest/test.git",
			"http_url": "https://gitlab.com/adwaittest/test.git"
		  },
		  "last_commit": {
			"id": "5b7fafc03361d1d021e8f9214879e16498630c5a",
			"message": "Add new file 5",
			"title": "Add new file 5",
			"timestamp": "2020-11-02T23:33:01+00:00",
			"url": "https://gitlab.com/adwaittest/test/-/commit/5b7fafc03361d1d021e8f9214879e16498630c5a",
			"author": {
			  "name": "Adwait Bhandare",
			  "email": "bhandare.adwait@gmail.com"
			}
		  },
		  "work_in_progress": false,
		  "total_time_spent": 0,
		  "human_total_time_spent": null,
		  "human_time_estimate": null,
		  "assignee_ids": [
	  
		  ],
		  "state": "opened",
		  "action": "open"
		},
		"labels": [
	  
		],
		"changes": {
		  "author_id": {
			"previous": null,
			"current": 7341710
		  },
		  "created_at": {
			"previous": null,
			"current": "2020-11-02 23:35:31 UTC"
		  },
		  "description": {
			"previous": null,
			"current": ""
		  },
		  "id": {
			"previous": null,
			"current": 76594814
		  },
		  "iid": {
			"previous": null,
			"current": 1
		  },
		  "merge_params": {
			"previous": {
			},
			"current": {
			  "force_remove_source_branch": "1"
			}
		  },
		  "source_branch": {
			"previous": null,
			"current": "stage"
		  },
		  "source_project_id": {
			"previous": null,
			"current": 21699507
		  },
		  "target_branch": {
			"previous": null,
			"current": "master"
		  },
		  "target_project_id": {
			"previous": null,
			"current": 21699507
		  },
		  "title": {
			"previous": null,
			"current": "Add new file 5"
		  },
		  "updated_at": {
			"previous": null,
			"current": "2020-11-02 23:35:31 UTC"
		  }
		},
		"repository": {
		  "name": "test",
		  "url": "git@gitlab.com:adwaittest/test.git",
		  "description": "",
		  "homepage": "https://gitlab.com/adwaittest/test"
		}
	  }`
}

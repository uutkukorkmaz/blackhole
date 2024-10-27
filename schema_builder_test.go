package blackhole

import (
	"testing"
)

func TestMySQLGrammar(t *testing.T) {
	schema := NewSchema(MySQL)

	schema.Create("builder_test", func(table *Blueprint) {
		table.Id()

		table.String("test_string_column", 255).
			NotNull().
			Unique()

		table.String("test_default", 255).
			Default("default_value").
			IndexUsing(IndexAlgorithmBTree)

		table.Enum("test_enum", []string{"a", "b", "c"}).
			Default("a")

		table.Timestamps()

		table.Collate("utf8mb4_unicode_ci")
	})

	expectedSQL := "create table if not exists `builder_test`(`id` bigint unsigned not null auto_increment primary key,`test_string_column` varchar(255) not null,`test_default` varchar(255) not null default 'default_value',`test_enum` enum('a','b','c') not null default 'a',`created_at` timestamp not null default CURRENT_TIMESTAMP,`updated_at` timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP) default character set utf8mb4 collate 'utf8mb4_unicode_ci';\nalter table `builder_test` add unique `builder_test_test_string_column_unique`(`test_string_column`);\nalter table `builder_test` add index `builder_test_test_default_index`(`test_default`) using btree;"
	generatedSQL, err := schema.Build()

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if generatedSQL != expectedSQL {
		t.Errorf("Expected: %s", expectedSQL)
		t.Errorf("Got: %s", generatedSQL)
	}
}

func TestSchema_Create_WithMySQLGrammar(t *testing.T) {
	var cases = []struct {
		name     string
		callback func(*Blueprint)
		expected string
	}{
		{
			name: "users",
			callback: func(bp *Blueprint) {
				bp.Id()
				bp.String("username", 255).NotNull().Unique()
				bp.String("password", 255).NotNull()
				bp.Int("age").IndexUsing(IndexAlgorithmBTree)
			},
			expected: "create table if not exists `users`(`id` bigint unsigned not null auto_increment primary key,`username` varchar(255) not null,`password` varchar(255) not null,`age` integer(11)) default character set utf8mb4 collate 'utf8mb4_unicode_ci';\nalter table `users` add unique `users_username_unique`(`username`);\nalter table `users` add index `users_age_index`(`age`) using btree;",
		},
		{
			name: "posts",
			callback: func(bp *Blueprint) {
				bp.Id()
				bp.String("title", 255).NotNull()
				foreign, userId := bp.ForeignId("user_id")
				foreign.CascadeOnDelete()
				userId.Nullable()
			},
			expected: "create table if not exists `posts`(`id` bigint unsigned not null auto_increment primary key,`title` varchar(255) not null,`user_id` bigint unsigned null) default character set utf8mb4 collate 'utf8mb4_unicode_ci';\nalter table `posts` add constraint `posts_user_id_foreign` foreign key (`user_id`) references `users` (`id`) on delete cascade;",
		},
		{
			name: "posts",
			callback: func(bp *Blueprint) {
				bp.Id()
				bp.String("title", 255).NotNull()
				foreign, userId := bp.ForeignId("user_id")
				foreign.CascadeOnDelete()
				userId.Nullable()
				foreign.On("users", "id")
			},
			expected: "create table if not exists `posts`(`id` bigint unsigned not null auto_increment primary key,`title` varchar(255) not null,`user_id` bigint unsigned null) default character set utf8mb4 collate 'utf8mb4_unicode_ci';\nalter table `posts` add constraint `posts_user_id_foreign` foreign key (`user_id`) references `users` (`id`) on delete cascade;",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			schema := NewSchema(MySQL)
			schema.Create(c.name, c.callback)
			generatedSQL, err := schema.Build()

			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if generatedSQL != c.expected {
				t.Errorf("Expected: %s", c.expected)
				t.Errorf("Got: %s", generatedSQL)
			}
		})
	}
}

func TestSchema_Alter_WithMySQLGrammar(t *testing.T) {
	var cases = []struct {
		name     string
		callback func(*Blueprint)
		expected string
	}{
		{
			name: "users",
			callback: func(bp *Blueprint) {
				bp.Timestamps()
				bp.RenameColumn("name", "full_name")
				bp.DropColumn("users", "email")
			},
			expected: "alter table `users` add `created_at` timestamp not null default CURRENT_TIMESTAMP;\nalter table `users` add `updated_at` timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;\nalter table `users` rename column `name` to `full_name`;\nalter table `users` drop column `email`;",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			schema := NewSchema(MySQL)
			schema.Alter(c.name, c.callback)
			generatedSQL, err := schema.Build()

			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if generatedSQL != c.expected {
				t.Errorf("Expected: %s", c.expected)
				t.Errorf("Got: %s", generatedSQL)
			}
		})
	}
}

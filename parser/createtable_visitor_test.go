/*
 * MIT License
 *
 * Copyright (c) 2021 zeromicro
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package parser

import "testing"

func Test_onlyTableName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"foo", "foo"},
		{"foo`.`bar", "bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onlyTableName(tt.name); got != tt.want {
				t.Errorf("onlyTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIfPrimaryKeyExists(t *testing.T) {
	type args struct {
		constraints []*TableConstraint
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No constraints (nil slice)",
			args: args{
				constraints: nil,
			},
			want: false,
		},
		{
			name: "Empty constraints slice",
			args: args{
				constraints: []*TableConstraint{},
			},
			want: false,
		},
		{
			name: "One constraint without primary key",
			args: args{
				constraints: []*TableConstraint{
					{
						ColumnPrimaryKey: []string{},
					},
				},
			},
			want: false,
		},
		{
			name: "Multiple constraints, none with primary key",
			args: args{
				constraints: []*TableConstraint{
					{ColumnPrimaryKey: []string{}},
					{ColumnPrimaryKey: []string{}},
				},
			},
			want: false,
		},
		{
			name: "One constraint with primary key",
			args: args{
				constraints: []*TableConstraint{
					{
						ColumnPrimaryKey: []string{"id"},
					},
				},
			},
			want: true,
		},
		{
			name: "Multiple constraints, one with primary key",
			args: args{
				constraints: []*TableConstraint{
					{ColumnPrimaryKey: []string{}},
					{ColumnPrimaryKey: []string{"user_id"}},
					{ColumnPrimaryKey: []string{}},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrimaryKeyExists(tt.args.constraints); got != tt.want {
				t.Errorf("checkIfPrimaryKeyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

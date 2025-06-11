
_default:
    @just --list


person_proto:
    @rm -rf gen/*
    @protoc --go_opt=paths=source_relative --go_out=. person.proto

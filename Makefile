.PHONY: proto

proto:
	    protoc --twirp_out=. --go_out=. unumpb/unumpb.proto

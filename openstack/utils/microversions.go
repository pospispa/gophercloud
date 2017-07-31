package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ErrIncompatible is an error in case Incompatible microversion occured.
var ErrIncompatible = errors.New("Incompatible microversion")

// CompatibleMicroversion checks input microversions to find out whether they are compatible or not. The input parameters are:
// - minimumMV: minimum microversion supported by the called function.
// - maximumMV: maximum microversion supported by the called function.
// - requestedMV: microversion requested by the user of the gophercloud library.
// - serverMaximumMV: maximum microversion supported by the particular server the gophercloud library is talking to.
func CompatibleMicroversion(minimumMV, maximumMV, requestedMV, serverMinimumMV, serverMaximumMV string) error {
	minimumMajor, minimumMinor := splitMicroversion(minimumMV)
	if requestedMV == "latest" {
		err := validMicroversion(serverMaximumMV)
		if err == nil {
			serverMaximumMajor, serverMaximumMinor := splitMicroversion(serverMaximumMV)
			if (minimumMajor > serverMaximumMajor) || (minimumMajor == serverMaximumMajor && minimumMinor > serverMaximumMinor) {
				return ErrIncompatible
			}
		}
		if maximumMV == "" {
			return nil
		}
		if err != nil {
			return err
		}
		serverMaximumMajor, serverMaximumMinor := splitMicroversion(serverMaximumMV)
		maximumMajor, maximumMinor := splitMicroversion(maximumMV)
		if (maximumMajor > serverMaximumMajor) || (maximumMajor == serverMaximumMajor && maximumMinor >= serverMaximumMinor) {
			return nil
		}
		return ErrIncompatible
	}
	if requestedMV == "" {
		if err := validMicroversion(serverMinimumMV); err != nil {
			return ErrIncompatible
		}
		serverMinimumMajor, serverMinimumMinor := splitMicroversion(serverMinimumMV)
		if (serverMinimumMajor > minimumMajor) || (serverMinimumMajor == minimumMajor && serverMinimumMinor >= minimumMinor) {
			return nil
		}
		return ErrIncompatible
	}
	if err := validMicroversion(requestedMV); err != nil {
		return err
	}
	requestedMajor, requestedMinor := splitMicroversion(requestedMV)
	if maximumMV == "" {
		if (requestedMajor > minimumMajor) || (requestedMajor == minimumMajor && requestedMinor >= minimumMinor) {
			return nil
		}
	} else {
		maximumMajor, maximumMinor := splitMicroversion(maximumMV)
		if (requestedMajor > minimumMajor) || (requestedMajor == minimumMajor && requestedMinor >= minimumMinor) && (maximumMajor > requestedMajor) || (maximumMajor == requestedMajor && maximumMinor >= requestedMinor) {
			return nil
		}
	}
	return ErrIncompatible
}

func splitMicroversion(mv string) (major, minor int) {
	if err := validMicroversion(mv); err != nil {
		return
	}

	mvParts := strings.Split(mv, ".")
	major, _ = strconv.Atoi(mvParts[0])
	minor, _ = strconv.Atoi(mvParts[1])

	return
}

func validMicroversion(mv string) (err error) {
	if mv == "latest" {
		return
	}

	mvRe := regexp.MustCompile("^\\d+\\.\\d+$")
	if v := mvRe.MatchString(mv); v {
		return
	}

	err = ErrIncompatible
	return
}

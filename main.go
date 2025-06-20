package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
    var unixFlag, utcFlag bool

    var rootCmd = &cobra.Command{
        Use:   "sudc [flags] <expression>",
        Short: "Simple Unix date calculator",
        Args:  cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            expression := strings.Join(args, " ")

            if unixFlag && utcFlag {
                fmt.Println("Error: cannot use both --unix and --utc flags together")
                os.Exit(1)
            }

            result, err := evaluateExpression(expression, unixFlag, utcFlag)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                os.Exit(1)
            }

            fmt.Println(result)
        },
    }

    rootCmd.Flags().BoolVar(&unixFlag, "unix", false, "output in Unix timestamp format")
    rootCmd.Flags().BoolVar(&utcFlag, "utc", false, "output in UTC format")

    rootCmd.SetHelpTemplate(`Usage: sudc [--unix|--utc] <expression>
Examples:
  sudc now --unix
  sudc --unix "now-2d"
  sudc --utc 1750071305-1749898505

{{.UsageString}}`)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func evaluateExpression(expr string, unixOutput, utcOutput bool) (string, error) {
	// Handle "now" cases
	if strings.HasPrefix(expr, "now") {
		t := time.Now()
		if expr == "now" {
			return formatOutput(t, unixOutput, utcOutput), nil
		}

		// Parse time modification (like "now-2d")
		duration, err := parseDuration(strings.TrimPrefix(expr, "now"))
		if err != nil {
			return "", err
		}
		modifiedTime := t.Add(duration)
		return formatOutput(modifiedTime, unixOutput, utcOutput), nil
	}

	// Handle Unix timestamp subtraction (like "1750071305-1749898505")
	if strings.Contains(expr, "-") && !strings.HasPrefix(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid expression format")
		}

		t1, err := parseUnixTime(parts[0])
		if err != nil {
			return "", err
		}

		t2, err := parseUnixTime(parts[1])
		if err != nil {
			return "", err
		}

		duration := t1.Sub(t2)
		return formatDuration(duration), nil
	}

	// Handle single Unix timestamp
	if unixOutput || utcOutput {
		t, err := parseUnixTime(expr)
		if err != nil {
			return "", err
		}
		return formatOutput(t, unixOutput, utcOutput), nil
	}

	// Default output format if no flags specified
	t, err := parseUnixTime(expr)
	if err == nil {
		return formatOutput(t, unixOutput, utcOutput), nil
	}

	return "", fmt.Errorf("invalid expression")
}

func parseDuration(s string) (time.Duration, error) {
    if s == "" {
        return 0, nil
    }

    unitMap := map[string]time.Duration{
        "s": time.Second,
        "m": time.Minute,
        "h": time.Hour,
        "d": 24 * time.Hour,
    }

    var sign int64 = 1
    if s[0] == '-' {
        sign = -1
        s = s[1:]
    } else if s[0] == '+' {
        s = s[1:]
    }

    var num int64
    var unit string

    n, err := fmt.Sscanf(s, "%d%s", &num, &unit)
    if err != nil || n != 2 {
        return 0, fmt.Errorf("invalid duration format")
    }

    durationUnit, ok := unitMap[unit]
    if !ok {
        return 0, fmt.Errorf("unknown duration unit: %s", unit)
    }

    return time.Duration(sign * num) * durationUnit, nil
}

func parseUnixTime(s string) (time.Time, error) {
	var unixTime int64
	_, err := fmt.Sscanf(s, "%d", &unixTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid Unix timestamp")
	}
	return time.Unix(unixTime, 0), nil
}

func formatOutput(t time.Time, unixOutput, utcOutput bool) string {
	if unixOutput {
		return fmt.Sprintf("%d", t.Unix())
	}
	if utcOutput {
		return t.UTC().Format(time.RFC3339)
	}
	return t.Format(time.RFC3339)
}

func formatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour

	hours := d / time.Hour
	d -= hours * time.Hour

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}
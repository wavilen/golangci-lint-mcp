<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Integration Test Report — Phase 77</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; max-width: 1200px; margin: 0 auto; padding: 20px; background: #f8fafc; }
        h1 { color: #0f172a; border-bottom: 2px solid #e2e8f0; padding-bottom: 8px; }
        .pass { color: #16a34a; font-weight: bold; }
        .fail { color: #dc2626; font-weight: bold; }
        .skip { color: #d97706; }
        .summary { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 16px; margin: 24px 0; }
        .card { background: white; border: 1px solid #e5e7eb; border-radius: 8px; padding: 16px; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
        .card h3 { margin: 0 0 8px 0; color: #64748b; font-size: 0.85em; text-transform: uppercase; }
        .card .value { font-size: 2em; font-weight: bold; color: #0f172a; }
        table { border-collapse: collapse; width: 100%; margin: 20px 0; background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
        th { background: #f1f5f9; color: #475569; text-align: left; padding: 10px 12px; font-size: 0.85em; text-transform: uppercase; }
        td { border-top: 1px solid #e5e7eb; padding: 10px 12px; }
        tr:hover td { background: #f8fafc; }
        .bar { height: 8px; border-radius: 4px; background: #e5e7eb; overflow: hidden; }
        .bar-fill { height: 100%; border-radius: 4px; background: #16a34a; }
    </style>
</head>
<body>
    <h1>Integration Test Report</h1>
    <p style="color:#64748b">Generated: {{.Timestamp.Format "2006-01-02 15:04:05"}}</p>

    <div class="summary">
        <div class="card">
            <h3>Pass Rate</h3>
            <div class="value">{{printf "%.0f" .PassRate}}%</div>
            <div style="color:#64748b;margin-top:4px">{{.PassCount}}/{{.TotalCount}} tests</div>
            <div class="bar" style="margin-top:8px"><div class="bar-fill" style="width:{{printf "%.0f" .PassRate}}%"></div></div>
        </div>
        <div class="card">
            <h3>Models Tested</h3>
            <div style="margin-top:8px">{{range .Models}}<div style="padding:2px 0">{{.}}</div>{{end}}</div>
        </div>
        <div class="card">
            <h3>Fixtures</h3>
            <div style="margin-top:8px">{{range .Fixtures}}<div style="padding:2px 0">{{.}}</div>{{end}}</div>
        </div>
        <div class="card">
            <h3>Avg Issue Reduction</h3>
            <div class="value">{{printf "%.0f" .AvgReduction}}%</div>
        </div>
    </div>

    <h2>Model Comparison</h2>
    <table>
        <tr>
            <th>Fixture</th>
            <th>Model</th>
            <th>Before</th>
            <th>After</th>
            <th>Reduction</th>
            <th>Builds</th>
            <th>nolint</th>
            <th>Tool Calls</th>
            <th>Retries</th>
            <th>Status</th>
        </tr>
        {{range .Results}}
        <tr>
            <td><strong>{{.TestCase}}</strong></td>
            <td>{{.Model}}</td>
            <td>{{.BeforeIssues}}</td>
            <td>{{.AfterIssues}}</td>
            <td>{{printf "%.1f" .IssueReduction}}%</td>
            <td>{{if .BuildsClean}}<span class="pass">✓</span>{{else}}<span class="fail">✗</span>{{end}}</td>
            <td>{{.NolintCount}}</td>
            <td>{{.ToolCalls}}</td>
            <td>{{.Retries}}</td>
            <td>{{if .Pass}}<span class="pass">PASS</span>{{else if .Error}}<span class="fail">ERROR</span>{{else}}<span class="skip">PARTIAL</span>{{end}}</td>
        </tr>
        {{end}}
    </table>

    {{if .Errors}}
    <h2>Errors</h2>
    <ul>
    {{range .Errors}}
        <li>{{.}}</li>
    {{end}}
    </ul>
    {{end}}
</body>
</html>

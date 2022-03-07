{{- range $key, $dep := .Deps -}}
import { {{ $dep }} } from './{{ $dep | ToKebabCase }}'
{{ end }}
export class {{ .Model.Name }} {
  {{- range .Model.Props}}
  {{ .Name }}: {{ .Type }}{{ if .IsArray }}[]{{ end }}{{ if HaveDefaultValue . }} = {{ PropDefaultValue . }}{{ end }}
  {{- end }}
}

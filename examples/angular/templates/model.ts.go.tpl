{{- range $key, $dep := ModelDeps .Model -}}
import { {{ $dep.Name }} } from './{{ $dep.Name | KebabCase }}'
{{ end }}
export class {{ .Model.Name }} {
  {{- range .Model.Props}}
  {{ .Name }}: {{ .Type }}{{ if .IsArray }}[]{{ end }}{{ if .DefaultValue }} = {{ .DefaultValue }}{{ end }}
  {{- end }}
}

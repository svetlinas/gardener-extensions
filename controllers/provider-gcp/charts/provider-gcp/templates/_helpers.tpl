{{- define "name" -}}
gardener-extension-provider-gcp
{{- end -}}

{{- define "labels.app.key" -}}
app.kubernetes.io/name
{{- end -}}
{{- define "labels.app.value" -}}
{{ include "name" . }}
{{- end -}}

{{- define "labels" -}}
{{ include "labels.app.key" . }}: {{ include "labels.app.value" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
helm.sh/chart: {{ include "name" . }}
{{- end -}}

{{-  define "image" -}}
  {{- if hasPrefix "sha256:" .Values.image.tag }}
  {{- printf "%s@%s" .Values.image.repository .Values.image.tag }}
  {{- else }}
  {{- printf "%s:%s" .Values.image.repository .Values.image.tag }}
  {{- end }}
{{- end }}
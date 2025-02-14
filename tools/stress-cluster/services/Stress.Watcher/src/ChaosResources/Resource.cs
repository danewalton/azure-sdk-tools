using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using k8s.Models;

namespace Stress.Watcher
{
    public class GenericChaosResource : CustomResource<ChaosResourceSpec, ChaosResourceStatus>
    {
        public static string PauseAnnotationKey = "experiment.chaos-mesh.org/pause";

        public static string TestInstanceLabelKey = "testInstance";

        public bool IsPaused()
        {
            var paused = "";
            Metadata.Annotations?.TryGetValue(GenericChaosResource.PauseAnnotationKey, out paused);
            return paused == "true";
        }

        public override string ToString()
        {
            var labels = "";
            foreach (var kvp in Metadata.Labels)
            {
                labels += kvp.Key + " : " + kvp.Value + ", ";
            }
            labels = labels.TrimEnd(',', ' ');
            return $"{Metadata.Name} Status: { CStatus.Experiment.phase }, Labels: {{ {labels} }}, Spec: {{ {Spec.Selector} }}";
        }
    }

    public class ChaosResourceSpec
    {
        [JsonProperty(PropertyName = "selector")]
        public ChaosSelector Selector { get; set; }
    }

    public class ChaosSelector
    {
        public ChaosLabelSelectors LabelSelectors { get; set; }
        public List<string> Namespaces { get; set; }

        public override string ToString()
        {
            return $"Namespaces: {String.Join(',', Namespaces)}, {LabelSelectors}";
        }
    }

    public class ChaosLabelSelectors
    {
        public string TestInstance { get; set; }

        public override string ToString()
        {
            return $"TestInstance: {TestInstance}";
        }
    }

    public class ChaosResourceStatus : V1Status
    {
        [JsonProperty(PropertyName = "experiment", NullValueHandling = NullValueHandling.Ignore)]
        public ChaosExperimentStatus Experiment { get; set; }
    }

    public class ChaosExperimentStatus
    {
        public string duration { get; set; }
        public DateTime endTime { get; set; }
        public string phase { get; set; }
        public DateTime startTime { get; set; }
    }
}

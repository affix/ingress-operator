package controller

import "testing"

func TestGetIssuerKind(t *testing.T) {
  issuer := getIssuerKind("ClusterIssuer")
  if(issuer != "certmanager.k8s.io/cluster-issuer") {
    t.Errorf("%s is not certmanager.k8s.io/cluster-issuer", issuer)
  }

  issuer = getIssuerKind("Issuer")
  if(issuer != "certmanager.k8s.io/issuer") {
    t.Errorf("%s is not certmanager.k8s.io/issuer", issuer)
  }

  issuer = getIssuerKind("NotaRealIssuer")
  if(issuer != "certmanager.k8s.io/issuer") {
    t.Errorf("%s is not certmanager.k8s.io/issuer", issuer)
  }
}

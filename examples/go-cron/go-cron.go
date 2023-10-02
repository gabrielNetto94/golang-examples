package gocron

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

type Protocol struct {
	WorkID int `json:"workId"`
	Acao   int `json:"acao"`
}

type Response struct {
	Code     int           `json:"code"`
	Debug    bool          `json:"debug"`
	WorkID   int           `json:"workId"`
	Tarefas  []interface{} `json:"tarefas"`
	WorkVid  int           `json:"workVid"`
	Processo struct {
		ProcIDPai   int    `json:"procIdPai"`
		ProcCodigo  int    `json:"procCodigo"`
		ObsAbertura string `json:"obsAbertura"`
		TprocessoID int    `json:"tprocessoId"`
	} `json:"processo"`
	Rascunho  bool `json:"rascunho"`
	Atributos []struct {
		Tipo     int           `json:"tipo"`
		Label    string        `json:"label"`
		Ordem    int           `json:"ordem"`
		TbAux    int           `json:"tbAux"`
		Value    string        `json:"value"`
		CalcID   string        `json:"calcId"`
		Codigo   string        `json:"codigo"`
		Oculto   bool          `json:"oculto"`
		Opcoes   []interface{} `json:"opcoes"`
		Regras   []interface{} `json:"regras"`
		AuxTipo  int           `json:"auxTipo"`
		Gatilho  []interface{} `json:"gatilho"`
		Mascara  string        `json:"mascara"`
		Tamanho  int           `json:"tamanho"`
		Visivel  bool          `json:"visivel"`
		Editavel bool          `json:"editavel"`
		Features struct {
			Maximo struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"maximo"`
			Minimo struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"minimo"`
			Indicador struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"indicador"`
			CopiaButton struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"copiaButton"`
			Informativo struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"informativo"`
			ValorPadrao struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"valorPadrao"`
			PesquisaProc struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"pesquisaProc"`
			HistoricoProc struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"historicoProc"`
			CopiaSemMascara struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"copiaSemMascara"`
			DadosIntegradora struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"dadosIntegradora"`
		} `json:"features,omitempty"`
		PformsID            int           `json:"pformsId"`
		TtableID            int           `json:"ttableId"`
		ConjDados           string        `json:"conjDados"`
		TipoTable           []interface{} `json:"tipoTable"`
		CampoChave          bool          `json:"campoChave"`
		PreSelecao          bool          `json:"preSelecao"`
		ValueTable          []interface{} `json:"valueTable"`
		AuxAtributo         string        `json:"auxAtributo"`
		CalcFormula         string        `json:"calcFormula"`
		InfoExterna         bool          `json:"infoExterna"`
		Integradora         bool          `json:"integradora"`
		Obrigatorio         bool          `json:"obrigatorio"`
		TypeRetorno         int           `json:"typeRetorno"`
		ValorPadrao         string        `json:"valorPadrao"`
		HistProcesso        bool          `json:"histProcesso"`
		MotivadoRegra       []int         `json:"motivadoRegra"`
		PrimeiraOpcao       string        `json:"primeiraOpcao"`
		LimiteCaracter      int           `json:"limiteCaracter"`
		CalcFormulaTable    interface{}   `json:"calcFormulaTable"`
		MotivaRegraAnexo    bool          `json:"motivaRegraAnexo"`
		TextoExplicativo    string        `json:"textoExplicativo"`
		MotivaRegraImpresso bool          `json:"motivaRegraImpresso"`
		Features0           struct {
			Mascara struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"mascara"`
			Indicador struct {
				Value   bool `json:"value"`
				Checked bool `json:"checked"`
			} `json:"indicador"`
			CopiaButton struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"copiaButton"`
			Informativo struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"informativo"`
			ValorPadrao struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"valorPadrao"`
			HistoricoDoc struct {
				Value   bool `json:"value"`
				Checked bool `json:"checked"`
			} `json:"historicoDoc"`
			PesquisaProc struct {
				Value   bool `json:"value"`
				Checked bool `json:"checked"`
			} `json:"pesquisaProc"`
			HistoricoProc struct {
				Value   bool `json:"value"`
				Checked bool `json:"checked"`
			} `json:"historicoProc"`
			SomenteNumero struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"somenteNumero"`
			LimiteCaracter struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"limiteCaracter"`
			CopiaSemMascara struct {
				Value   interface{} `json:"value"`
				Checked bool        `json:"checked"`
			} `json:"copiaSemMascara"`
			DadosIntegradora struct {
				Value   bool `json:"value"`
				Checked bool `json:"checked"`
			} `json:"dadosIntegradora"`
		} `json:"features,omitempty"`
	} `json:"atributos"`
	Confeccao           bool          `json:"confeccao"`
	Protocolo           interface{}   `json:"protocolo"`
	WorkLabel           string        `json:"workLabel"`
	TempoLabel          string        `json:"tempoLabel"`
	ProcessoOcr         []interface{} `json:"processoOcr"`
	VersaoArvore        string        `json:"versaoArvore"`
	TprocessoLabel      string        `json:"tprocessoLabel"`
	WorkExigeParecer    int           `json:"workExigeParecer"`
	WorkFormAgrupado    bool          `json:"workFormAgrupado"`
	AcoesIntegradoras   []interface{} `json:"acoesIntegradoras"`
	AtributosForaNodo   []interface{} `json:"atributosForaNodo"`
	WorkAnteriorLabel   string        `json:"workAnteriorLabel"`
	AssinaturaPendente  bool          `json:"assinaturaPendente"`
	WorkProtocoloRapido bool          `json:"workProtocoloRapido"`
}

func task() {
	fmt.Println("I am running task.")
}
func GoCron() {

	r := gin.Default()

	s := gocron.NewScheduler(time.UTC)
	r.GET("/ping", func(c *gin.Context) {

		s.Stop()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/start", func(c *gin.Context) {

		s.StartAsync()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s.Every(10).Seconds().Do(func() {
		fmt.Println("job 10")
	})

	s.Every(5).Seconds().Do(func() {
		fmt.Println("job 5")
	})

	s.Every(2).Seconds().Do(func() {
		fmt.Println("job 2")
	})
	s.StartAsync()

	r.Run() // listen
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

type Teste interface {
	int | float64
}

func Hello[T Teste](val T) {
	fmt.Println("hello: ", val)
}

func DefaultArgs(args ...int) {

}

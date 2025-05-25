class ReportBuild:
    def build(self):
        # logic to build report 
        pass

class ReportSave:  
    def save_to_file(self):
        # logic to save report to a pdf/xml/xlsx/word file

if __name__ == '__main__':
    rb = ReportBuild()
    rb.build()

    rs = ReportSave()
    rs.save_to_file

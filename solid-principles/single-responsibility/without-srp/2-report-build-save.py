class ReportBuildSave:
    def build(self):
        # logic to build report 
        pass
    
    def save_to_file(self):
        # logic to save report to a pdf/xml/xlsx/word file

if __name__ == '__main__':
    r = ReportBuildSave()
    r.build()
    r.save_to_file
